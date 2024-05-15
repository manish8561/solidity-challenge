import {
  time,
  loadFixture,
} from "@nomicfoundation/hardhat-toolbox/network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import hre from "hardhat";

const SECOND_IN_HOUR = 60 * 60;
const SECOND_IN_10_MIN = 10 * 60;
describe("BiddingWar", () => {
  // We define a fixture to reuse the same setup in every test.
  // We use loadFixture to run this setup once, snapshot that state,
  // and reset Hardhat Network to that snapshot in every test.
  async function deployBiddingWarFixture() {
    const BiddingWar = await hre.ethers.getContractFactory("BiddingWar");
    const biddingWar = await BiddingWar.deploy();

    // Contracts are deployed using the first signer/account by default
    const [owner, otherAccount] = await hre.ethers.getSigners();

    return { biddingWar, owner, otherAccount };
  }

  describe("Deployment", () => {
    it("Should set the right owner", async () => {
      const { biddingWar, owner } = await loadFixture(deployBiddingWarFixture);
      expect(await biddingWar.owner()).to.equal(owner.address);
    });

    it("Start the first game", async () => {
      const { biddingWar } = await loadFixture(deployBiddingWarFixture);
      const gameStatus = await biddingWar.gameStatus();
      expect(gameStatus).to.equal(true);
    });
    it("Restart the second game", async () => {
      const { biddingWar } = await loadFixture(deployBiddingWarFixture);
      const gameStatus = await biddingWar.gameStatus();
      expect(gameStatus).to.equal(true);

      await expect(biddingWar.declareResult()).to.be.revertedWith(
        "Last game is still running!"
      );
    });
    it("Restart the second game when previous game time expired", async () => {
      const { biddingWar } = await loadFixture(deployBiddingWarFixture);
      const gameStatus = await biddingWar.gameStatus();
      expect(gameStatus).to.equal(true);

      // We can increase the time in Hardhat Network
      const afterOneHour = (await time.latest()) + SECOND_IN_HOUR;
      await time.increaseTo(afterOneHour);

      await expect(biddingWar.declareResult()).not.to.be.reverted;
    });
  });

  describe("Place the bid and declare result", () => {
    describe("Place the bids", () => {
      it("Should place the bid with zero value", async () => {
        const { biddingWar, owner, otherAccount } = await loadFixture(
          deployBiddingWarFixture
        );

        await expect(
          biddingWar.connect(otherAccount).placeBid({ value: 0 })
        ).to.be.revertedWith("Bid amount should be greater than zero!");
      });
      it("Should place the first bid", async () => {
        const { biddingWar, owner, otherAccount } = await loadFixture(
          deployBiddingWarFixture
        );

        await biddingWar.connect(otherAccount).placeBid({ value: 100 });
        expect(await biddingWar.bidCounter()).to.equal(1);

        expect(
          await hre.ethers.provider.getBalance(biddingWar.target)
        ).to.equal(95);
      });
      it("Should place the second bid with same amount", async () => {
        const { biddingWar, owner, otherAccount } = await loadFixture(
          deployBiddingWarFixture
        );

        await expect(biddingWar.connect(otherAccount).placeBid({ value: 100 }))
          .not.to.be.reverted;
        expect(await biddingWar.bidCounter()).to.equal(1);

        expect(
          await hre.ethers.provider.getBalance(biddingWar.target)
        ).to.equal(95);

        await expect(
          biddingWar.connect(otherAccount).placeBid({ value: 100 })
        ).to.be.revertedWith(
          "Bid amount should be greater than last bid amount!"
        );
      });

      it("Should place the bid with after game endTime", async () => {
        const { biddingWar, owner, otherAccount } = await loadFixture(
          deployBiddingWarFixture
        );

        // We can increase the time in Hardhat Network
        const afterOneHour = (await time.latest()) + SECOND_IN_HOUR;
        await time.increaseTo(afterOneHour);

        await expect(
          biddingWar.connect(otherAccount).placeBid({ value: 100 })
        ).to.be.revertedWith("Game over wait for new game!");
      });

      it("Should place the second bid with higher amount", async () => {
        const { biddingWar, owner, otherAccount } = await loadFixture(
          deployBiddingWarFixture
        );

        await expect(biddingWar.connect(otherAccount).placeBid({ value: 100 }))
          .not.to.be.reverted;
        expect(await biddingWar.bidCounter()).to.equal(1);

        expect(
          await hre.ethers.provider.getBalance(biddingWar.target)
        ).to.equal(95);

        await expect(biddingWar.connect(otherAccount).placeBid({ value: 120 }))
          .not.to.be.reverted;
        expect(
          await hre.ethers.provider.getBalance(biddingWar.target)
        ).to.equal(209);
      });
    });
    describe("Declare result", () => {
      it("Should declare result before game expire.", async () => {
        const { biddingWar, owner, otherAccount } = await loadFixture(
          deployBiddingWarFixture
        );
        await expect(biddingWar.declareResult()).to.be.revertedWith(
          "Last game is still running!"
        );

        // We can increase the time in Hardhat Network
        const afterOneHour = (await time.latest()) + SECOND_IN_HOUR;
        await time.increaseTo(afterOneHour);
        await expect(biddingWar.declareResult()).not.to.be.reverted;
      });
      it("Should declare the result", async () => {
        const { biddingWar, owner, otherAccount } = await loadFixture(
          deployBiddingWarFixture
        );

        await expect(biddingWar.connect(otherAccount).placeBid({ value: 100 }))
          .not.to.be.reverted;
        expect(await biddingWar.bidCounter()).to.equal(1);

        expect(
          await hre.ethers.provider.getBalance(biddingWar.target)
        ).to.equal(95);

        // We can increase the time in Hardhat Network
        const afterOneHour = (await time.latest()) + SECOND_IN_HOUR;
        await time.increaseTo(afterOneHour);

        await expect(biddingWar.connect(otherAccount).placeBid({ value: 120 }))
          .not.to.be.reverted;

        expect(
          await hre.ethers.provider.getBalance(biddingWar.target)
        ).to.equal(209);

        // We can increase the time in Hardhat Network
        const after10minutes = (await time.latest()) + SECOND_IN_10_MIN * 2;
        await time.increaseTo(after10minutes);

        await expect(biddingWar.declareResult()).not.to.be.reverted;

        expect(
          await hre.ethers.provider.getBalance(biddingWar.target)
        ).to.equal(0);
      });
    });
    describe("Events", () => {
      it("Should emit an event on place bid and declare result", async () => {
        const { biddingWar, owner, otherAccount } = await loadFixture(
          deployBiddingWarFixture
        );
        const value = 100;
        const winAmt = 95;
        let latestTime = (await time.latest()) + 1;
        await expect(biddingWar.connect(otherAccount).placeBid({ value }))
          .to.emit(biddingWar, "BidEvent")
          .withArgs(1, value, latestTime, otherAccount.address); // We accept any value as `when` arg

        // We can increase the time in Hardhat Network
        const afterOneHour = (await time.latest()) + SECOND_IN_HOUR;
        await time.increaseTo(afterOneHour);

        const after10minutes = (await time.latest()) + SECOND_IN_10_MIN;
        await time.increaseTo(after10minutes);

        latestTime = (await time.latest()) + 1;
        await expect(biddingWar.declareResult())
          .to.emit(biddingWar, "WinEvent")
          .withArgs(1, value, winAmt, latestTime, otherAccount.address);
        expect(
          await hre.ethers.provider.getBalance(biddingWar.target)
        ).to.equal(0);
      });
    });
    describe("Sending ether to the contract direcetly", () => {
      it("Should run the receive function", async () => {
        const { biddingWar, owner, otherAccount } = await loadFixture(
          deployBiddingWarFixture
        );
        // sending ethers directly to the contract address
        const value = 100;
        const data = {
          value,
          to: biddingWar.target
        }
        await expect(owner.sendTransaction(data)).not.to.be.reverted;

        expect(
          await hre.ethers.provider.getBalance(biddingWar.target)
        ).to.equal(0);
      });
    });
  });
});
