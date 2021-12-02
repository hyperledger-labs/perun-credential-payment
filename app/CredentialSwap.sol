// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import "./perun-eth-contracts/contracts/App.sol";
import "./perun-eth-contracts/contracts/Channel.sol";
import "./perun-eth-contracts/vendor/openzeppelin-contracts/contracts/cryptography/ECDSA.sol";
import "./perun-eth-contracts/contracts/Array.sol";
import "./Decode.sol";

/**
 * CredentialSwap is a channel app for swapping a credential against a payment.
 */
contract CredentialSwap is App {
    enum Mode{ Default, Offer, Cert }
    uint8 constant ASSET_INDEX = 0;
    // Indices corresponding to data encoding.
    uint8 constant MODE_INDEX = 0;
    uint8 constant SIG_INDEX = 0;
    uint8 constant SIG_LENGTH = 65;

    struct Frame {
        uint8 mode;
        bytes body;
    }

    struct Offer {
        address issuer;
        bytes32 h;
        uint256 price;
        uint16 buyer;
    }

    struct Cert {
        bytes sig;
    }

    /**
     * ValidTransition checks if the transition from `cur` to `next` by
     * participant `actor` is valid.
     *
     * @param cur The current state.
     * @param next The potential next state.
     * @param actor The index of the actor.
     */
    function validTransition(
        Channel.Params calldata /*params*/,
        Channel.State calldata cur,
        Channel.State calldata next,
        uint256 actor
    ) external pure override {
        // We require that we only have a constant single asset.
        requireConstantSingleAsset(cur, next);

        // Decode current state.
        Frame memory frame = decodeFrame(cur);
        if (frame.mode == uint8(Mode.Offer)) {
            Offer memory offer = decodeOffer(frame.body);
            validTransitionFromOffer(offer, cur, next, actor);
        } else {
            // We require that the balances did not change.
            requireBalancesUnchanged(cur, next);

            // If the next state is an offer, check that the potential buyer has
            // sufficient funds to fulfill the payment.
            Frame memory nextFrame = decodeFrame(next);
            if (nextFrame.mode == uint8(Mode.Offer)) {
                Offer memory offer = decodeOffer(nextFrame.body);
                uint256[][] calldata nextBals = next.outcome.balances;
                require(nextBals[ASSET_INDEX][offer.buyer] >= offer.price,
                    "insufficient funds");
            }
        }
    }

    function validTransitionFromOffer(
        Offer memory offer,
        Channel.State calldata cur,
        Channel.State calldata next,
        uint256 actor
    ) internal pure {
        // Decode next state.
        (Cert memory cert, bool ok) = decodeCert(next);
        require(ok, "invalid next mode");
        uint256 seller = actor;

        // Verify signature.
        require(verify(offer.h, cert.sig, offer.issuer), "invalid signature");

        // Verify balances.
        uint256[][] calldata curBals = cur.outcome.balances;
        uint256[][] calldata nextBals = next.outcome.balances;
        require(nextBals[ASSET_INDEX][offer.buyer] == curBals[ASSET_INDEX][offer.buyer] - offer.price,
            "invalid amount transferred: buyer");
        require(nextBals[ASSET_INDEX][seller] == curBals[ASSET_INDEX][seller] + offer.price,
            "invalid amount transferred: seller");
    }

    function decodeFrame(Channel.State calldata s) internal pure returns (Frame memory) {
        uint8 dataIndex = 2; // Length is encoded as uint16 at index 0. Data starts afterwards at index 2. Encoding the length is currently needed as the encoding is also used for our stream-based off-chain communication.
        uint256 length = s.appData.length - dataIndex;
        bytes memory data = Decode.slice(s.appData, dataIndex, length);
        (Frame memory frame) = abi.decode(data, (Frame));
        return frame;
    }

    function decodeOffer(bytes memory data) internal pure returns (Offer memory) {
        (Offer memory offer) = abi.decode(data, (Offer));
        return offer;
    }

    function decodeCert(Channel.State calldata state) internal pure returns (Cert memory, bool) {
        Frame memory s = decodeFrame(state);
        if (s.mode != uint8(Mode.Cert)) {
            return (Cert(''), false);
        }

        return (Cert({sig: s.body}), true);
    }

    /// verify verifies that `sig` is a signature on `h` by `signer`.
    function verify(bytes32 h, bytes memory sig, address signer) internal pure returns (bool) {
        address recoveredAddr = ECDSA.recover(h, sig);
        return recoveredAddr == signer;
    }

    function requireConstantSingleAsset(Channel.State calldata cur, Channel.State calldata next) internal pure {
        (address[] memory a1, address[] memory a2) = (cur.outcome.assets, next.outcome.assets);
        
        uint8 numAssets = 1;
        require(a1.length == numAssets, "invalid number of assets: current");
        require(a2.length == numAssets, "invalid number of assets: next");

        require(a1[ASSET_INDEX] == a2[ASSET_INDEX], "invalid asset");
    }

    function requireBalancesUnchanged(Channel.State calldata cur, Channel.State calldata next) internal pure {
        requireEqualUint256ArrayArray(cur.outcome.balances, next.outcome.balances);
    }

    function requireEqualUint256ArrayArray(uint256[][] memory a, uint256[][] memory b) internal pure {
        require(a.length == b.length, "uint256[][]: unequal length");
        for (uint i = 0; i < a.length; i++) {
            Array.requireEqualUint256Array(a[i], b[i]);
        }
    }
}
