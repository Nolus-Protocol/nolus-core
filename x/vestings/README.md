# Allocation

This module handles distributing inflation to specific modules and accounts.

It is currently used for distributing inflation to the NFT incentives module and to the developer rewards.

## Governance Parameters

Here's a table for the parameters of the module:

| Name                       | Type                      | Description                                            | Default   |
| -------------------------- | ------------------------- | ------------------------------------------------------ | --------- |
| `DistributionProportions`  | `DistributionProportions` | Allocations for `NftIncentives` and `DeveloperRewards` | See below |
| `DeveloperRewardsReceiver` | Array<WeightedAddress>                     | List of addresses with allocation weight               | `[]`      |

The `DistributionProportions` is a struct with the following fields:

| Name               | Type         | Description                                                 | Default  |
| ------------------ | ------------ | ----------------------------------------------------------- | -------- |
| `Address`    | string (dec) | Proportion of inflation to be distributed to NFT incentives | `"0.45"` |
| `DeveloperRewards` | string (dec) | Proportion of inflation to be distributed to developers     | `"0.15"` |


The `WeightedAddress` is a struct with the following fields:
| Name               | Type         | Description                                    | 
| ------------------ | ------------ | -----------------------------------------------|
| `Address`          | string       | Receiving address                              |
| `Weight`           | string (dec) | Proportion of the assigned developer rewards   |

