## config
## PUT
### 
Save/update configuration
Used by the config upload CLI
### Expected Response Types
| Response | Reason                |
| -------- | --------------------- |
| 201      |                       |
| 500      | Internal Server Error |

### Parameters
| Name                        | In   | Description                                 | Required? | Type   |
| --------------------------- | ---- | ------------------------------------------- | --------- | ------ |
| tracker.cxApplicationConfig | body | Request for creating/updating configuration | true      | object |

## config/
## GET
### 
Returns configuration for a given genesis Hash
Used by the auto-setup CX chain CLI
### Expected Response Types
| Response | Reason                |
| -------- | --------------------- |
|          |                       |
| 200      | OK                    |
| 404      | Not Found             |
| 500      | Internal Server Error |

### Parameters
| Name        | In    | Description        | Required? | Type   |
| ----------- | ----- | ------------------ | --------- | ------ |
| genesisHash | query | Config genesisHash | true      | string |

### Content Types Produced
| Produces         |
| ---------------- |
| application/json |

## configs
## GET
### 
Returns list of all stored configurations
Used by the web app to display a list of configs and servers
### Expected Response Types
| Response | Reason                |
| -------- | --------------------- |
| 200      | OK                    |
| 500      | Internal Server Error |

### Content Types Produced
| Produces         |
| ---------------- |
| application/json |

## Definitions
### api.ErrorResponse Definition
| Property | Type   | Format |
| -------- | ------ | ------ |
|          |        |        |
| message  | string |        |
### tracker.CxApplication Definition
| Property  | Type   | Format |
| --------- | ------ | ------ |
|           |        |        |
| chainType | string |        |
| config    | string |        |
| createdAt | string |        |
| servers   | array  |        |
| updatedAt | string |        |
### tracker.Server Definition
| Property  | Type   | Format |
| --------- | ------ | ------ |
|           |        |        |
| address   | string |        |
| createdAt | string |        |
| updatedAt | string |        |
