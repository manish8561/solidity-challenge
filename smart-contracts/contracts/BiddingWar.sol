// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

// Import Ownable from the OpenZeppelin Contracts library
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

// Uncomment this line to use console.log
// import "hardhat/console.sol";

contract BiddingWar is Ownable, ReentrancyGuard {
    uint public commission = 5;
    uint128 public bidCounter = 0;
    uint64 public gameEndTime;
    uint public highestBid;
    address public highestBidder;
    bool public gameStatus = false;

    struct Bid {
        uint amount;
        address user;
        uint64 when;
    }

    mapping(uint => Bid) public bids;

    event BidEvent(
        uint128 indexed bid,
        uint amount,
        uint64 when,
        address indexed user
    );

    event WinEvent(
        uint128 indexed bid,
        uint highestBid,
        uint winAmount,
        uint64 when,
        address indexed winner
    );

    constructor() Ownable(msg.sender) {
        startGame();
    }

    /**
     * start the game
     */
    function startGame() private {
        require(gameEndTime < block.timestamp, "Last game is still running!");
        require(!gameStatus, "Last game is still waiting for result!");

        gameEndTime = uint64(block.timestamp + 3600); //adding hour
        gameStatus = true; // start/restart the game status
        highestBidder = address(0);
        highestBid = 0;
    }

    /**
     * place the bid
     */
    function placeBid() external payable nonReentrant {
        require(msg.value > 0, "Bid amount should be greater than zero!");
        require(
            msg.value > highestBid,
            "Bid amount should be greater than last bid amount!"
        );
        require(gameEndTime > block.timestamp, "Game over wait for new game!");
        // require(gameStatus, "Wait for new game!");
        bidCounter++;
        bids[bidCounter].amount = msg.value;
        bids[bidCounter].when = uint64(block.timestamp);
        bids[bidCounter].user = msg.sender;

        emit BidEvent(
            bidCounter,
            msg.value,
            uint64(block.timestamp),
            msg.sender
        );
        // changing values in the game
        gameEndTime = gameEndTime + 600; // increase by 10 minutes
        highestBid = msg.value;
        highestBidder = msg.sender;

        //sending the commission
        uint commissionValue = (msg.value * commission) / 100;
        // payable(owner()).transfer(commissionValue);
        // recommended method to sent ether
        (bool sent, ) = payable(owner()).call{value: commissionValue}("");
        require(sent, "Failed to send Ether");
    }

    /**
     * declare result and restart the game
     */
    function declareResult() public onlyOwner {
        require(gameEndTime < block.timestamp, "Last game is still running!");
        gameStatus = false;
        //send funds only highet bidder have address
        if (highestBidder != address(0)) {
            // uint winAmt = address(this).balance;
            uint winAmt = 0;
            assembly {
                winAmt := balance(address())
            }
            // payable(highestBidder).transfer(winAmt);
            // recommended method to sent ether
            (bool sent,) = payable(highestBidder).call{value: winAmt}("");
            require(sent, "Failed to send Ether");
            emit WinEvent(
                bidCounter,
                highestBid,
                winAmt,
                uint64(block.timestamp),
                highestBidder
            );
        }
        // restart the games
        startGame();
    }

    /**
     * common function to transfer fund to owner
     */
    function fundsTransferToOwner() private {
        payable(owner()).transfer(msg.value);
    }

    receive() external payable {
        fundsTransferToOwner();
    }

    fallback() external payable {
        fundsTransferToOwner();
    }
}
