import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const BiddingWarModule = buildModule("BiddingWarModule", (m) => {
  const biddingWar = m.contract("BiddingWar");

  return { biddingWar };
});

export default BiddingWarModule;
