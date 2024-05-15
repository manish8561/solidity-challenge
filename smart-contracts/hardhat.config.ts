import { HardhatUserConfig, task, vars } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@nomicfoundation/hardhat-verify";
import "hardhat-gas-reporter";

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.24",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
      evmVersion:"istanbul",
    },
  },
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",
      accounts: {
        mnemonic: vars.get("MNEMONIC"),
        count: 20,
      },
    },
    energi: {
      url: "https://nodeapi.test.energi.network/v1/jsonrpc",
      accounts: {
        mnemonic: vars.get("MNEMONIC"),
        count: 5,
      },
    },
    hardhat: {
      // forking: {
      //   url: "https://nodeapi.test.energi.network/v1/jsonrpc",
      //   blockNumber: 2222569,
      // },
      accounts: {
        mnemonic: vars.get("MNEMONIC"),
        count: 20,
      },
    },
  },
  sourcify: {
    enabled: true,
  },
  gasReporter: {
    enabled: false,
  },
};

task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

task("balance", "Prints the list of accounts", async (taskArgs, hre) => {
  const [owner, otherAccount] = await hre.ethers.getSigners();
  const balance = await hre.ethers.provider.getBalance(owner.address);
  console.table({
    owner: owner.address,
    balance,
  });
});

export default config;
