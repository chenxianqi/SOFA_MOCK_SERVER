package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"],
        beego.ControllerComments{
            Method: "GetMerchantAPITokenStatus",
            Router: "/merchant/:merchant_id/apisecret",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"],
        beego.ControllerComments{
            Method: "ActivateMerchantAPIToken",
            Router: "/merchant/:merchant_id/apisecret/activate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"],
        beego.ControllerComments{
            Method: "RefreshMerchantSecret",
            Router: "/merchant/:merchant_id/apisecret/refreshsecret",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"],
        beego.ControllerComments{
            Method: "SetAPIToken",
            Router: "/merchant/:merchant_id/apitoken",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"],
        beego.ControllerComments{
            Method: "ResendFailedMerchantCallbacks",
            Router: "/merchant/:merchant_id/notifications/manual",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"],
        beego.ControllerComments{
            Method: "CancelPaymentOrder",
            Router: "/merchant/:merchant_id/order",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"],
        beego.ControllerComments{
            Method: "RequestPaymentOrder",
            Router: "/merchant/:merchant_id/order",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"],
        beego.ControllerComments{
            Method: "QueryPaymentOrder",
            Router: "/merchant/:merchant_id/order",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"],
        beego.ControllerComments{
            Method: "UpdateOrderDuration",
            Router: "/merchant/:merchant_id/order/duration",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:MerchantController"],
        beego.ControllerComments{
            Method: "Callback",
            Router: "/merchant/callback",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "CreateDepositWalletAddresses",
            Router: "/wallets/:wallet_id/addresses",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetDepositWalletAddresses",
            Router: "/wallets/:wallet_id/addresses",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetDeployedContractCollectionAddresses",
            Router: "/wallets/:wallet_id/addresses/contract_txid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetDepositWalletAddressesLabel",
            Router: "/wallets/:wallet_id/addresses/get_labels",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetInvalidDepositAddresses",
            Router: "/wallets/:wallet_id/addresses/invalid-deposit",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "UpdateDepositWalletAddressLabel",
            Router: "/wallets/:wallet_id/addresses/label",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "VerifyAddresses",
            Router: "/wallets/:wallet_id/addresses/verify",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetTxAPITokenStatus",
            Router: "/wallets/:wallet_id/apisecret",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "ActivateAPIToken",
            Router: "/wallets/:wallet_id/apisecret/activate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "SetAPIToken",
            Router: "/wallets/:wallet_id/apitoken",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetAutoFee",
            Router: "/wallets/:wallet_id/autofee",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetAutoFees",
            Router: "/wallets/:wallet_id/autofees",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetWalletBlockInfo",
            Router: "/wallets/:wallet_id/blocks",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "ResendDepositCollectionCallbacks",
            Router: "/wallets/:wallet_id/collection/notifications/manual",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "CallContractRead",
            Router: "/wallets/:wallet_id/contract/read",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetWalletInfo",
            Router: "/wallets/:wallet_id/info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetNotifications",
            Router: "/wallets/:wallet_id/notifications",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetCallbackBySerial",
            Router: "/wallets/:wallet_id/notifications/get_by_id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "NotificationsInspect",
            Router: "/wallets/:wallet_id/notifications/inspect",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetDepositWalletPoolAddress",
            Router: "/wallets/:wallet_id/pooladdress",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetDepositWalletPoolAddressBalance",
            Router: "/wallets/:wallet_id/pooladdress/balance",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "VerifyDepositAddresses",
            Router: "/wallets/:wallet_id/receiver/addresses/verify",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetDepositWalletBalance",
            Router: "/wallets/:wallet_id/receiver/balance",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetDepositCallback",
            Router: "/wallets/:wallet_id/receiver/notifications/txid/:txid/:vout_index",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "RefreshSecret",
            Router: "/wallets/:wallet_id/refreshsecret",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetWithdrawalWalletBalance",
            Router: "/wallets/:wallet_id/sender/balance",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "ResendWithdrawalCallbacks",
            Router: "/wallets/:wallet_id/sender/notifications/manual",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetWithdrawalCallback",
            Router: "/wallets/:wallet_id/sender/notifications/order_id/:order_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "WithdrawAssets",
            Router: "/wallets/:wallet_id/sender/transactions",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetSenderTransactionHistory",
            Router: "/wallets/:wallet_id/sender/transactions",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetWithdrawTransactionState",
            Router: "/wallets/:wallet_id/sender/transactions/:order_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetWithdrawTransactionStateAll",
            Router: "/wallets/:wallet_id/sender/transactions/:order_id/all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "CancelWithdrawTransactions",
            Router: "/wallets/:wallet_id/sender/transactions/:order_id/cancel",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "SetWithdrawalACL",
            Router: "/wallets/:wallet_id/sender/transactions/acl",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetTransactionEventLog",
            Router: "/wallets/:wallet_id/sender/transactions/eventlog",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "RemoveSenderWhitelist",
            Router: "/wallets/:wallet_id/sender/whitelist",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "AddSenderWhitelist",
            Router: "/wallets/:wallet_id/sender/whitelist",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetSenderWhitelist",
            Router: "/wallets/:wallet_id/sender/whitelist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "CheckSenderWhitelist",
            Router: "/wallets/:wallet_id/sender/whitelist/check",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "QuerySenderWhitelistConfig",
            Router: "/wallets/:wallet_id/sender/whitelist/config",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "SignMessage",
            Router: "/wallets/:wallet_id/signmessage",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetTransactionHistory",
            Router: "/wallets/:wallet_id/transactions",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetVaultWalletBalance",
            Router: "/wallets/:wallet_id/vault/balance",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "Callback",
            Router: "/wallets/callback",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetReadOnlyWalletList",
            Router: "/wallets/readonly/walletlist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "GetReadOnlyWalletListBalances",
            Router: "/wallets/readonly/walletlist/balances",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"] = append(beego.GlobalControllerRouter["SOFA_MOCK_SERVER/controllers:OuterController"],
        beego.ControllerComments{
            Method: "WithdrawalCallback",
            Router: "/wallets/withdrawal/callback",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
