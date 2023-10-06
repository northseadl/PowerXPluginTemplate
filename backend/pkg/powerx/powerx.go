package powerx

import (
	"PluginTemplate/pkg/powerx/client"
)

type PowerX struct {
	*client.PClient

	AdminCommon *AdminCommon

	AdminDepartment *AdminDepartment

	AdminEmployee *AdminEmployee

	AdminPermission *AdminPermission

	AdminAuth *AdminAuth

	AdminDictionary *AdminDictionary

	AdminUserinfo *AdminUserinfo

	AdminMediaresource *AdminMediaresource

	AdminCustomerdomainLeader *AdminCustomerdomainLeader

	AdminCustomerdomainCustomer *AdminCustomerdomainCustomer

	AdminMarketMedia *AdminMarketMedia

	AdminMarketStore *AdminMarketStore

	AdminBusinessOpportunity *AdminBusinessOpportunity

	AdminProductPricebook *AdminProductPricebook

	AdminProduct *AdminProduct

	AdminProductCategory *AdminProductCategory

	AdminProductProductspecific *AdminProductProductspecific

	AdminProductSku *AdminProductSku

	AdminProductPricebookentry *AdminProductPricebookentry

	AdminProductArtisan *AdminProductArtisan

	AdminTradeAddressShipping *AdminTradeAddressShipping

	AdminTradeAddressBilling *AdminTradeAddressBilling

	AdminTradeAddressDelivery *AdminTradeAddressDelivery

	AdminTradeWarehouse *AdminTradeWarehouse

	AdminContractway *AdminContractway

	AdminScrmCustomer *AdminScrmCustomer

	AdminScrmContact *AdminScrmContact

	MpDictionary *MpDictionary

	MpCustomerAuth *MpCustomerAuth

	MpMarketStore *MpMarketStore

	MpMarketMedia *MpMarketMedia

	MpProductArtisan *MpProductArtisan

	MpProduct *MpProduct

	MpTradeCart *MpTradeCart

	MpTradeOrder *MpTradeOrder

	AdminTradeOrder *AdminTradeOrder

	AdminTradePayment *AdminTradePayment

	MpTradeAddressShipping *MpTradeAddressShipping

	MpTradeAddressDelivery *MpTradeAddressDelivery

	MpTradeAddressBilling *MpTradeAddressBilling

	MpTradePayment *MpTradePayment

	WebCustomerAuth *WebCustomerAuth

	WebCustomerAuthOa *WebCustomerAuthOa
}

func NewPowerX(endpoint string, debug bool) *PowerX {
	power := &PowerX{
		PClient: client.NewPClient(endpoint, debug),
	}

	power.AdminCommon = &AdminCommon{power}

	power.AdminDepartment = &AdminDepartment{power}

	power.AdminEmployee = &AdminEmployee{power}

	power.AdminPermission = &AdminPermission{power}

	power.AdminAuth = &AdminAuth{power}

	power.AdminDictionary = &AdminDictionary{power}

	power.AdminUserinfo = &AdminUserinfo{power}

	power.AdminMediaresource = &AdminMediaresource{power}

	power.AdminCustomerdomainLeader = &AdminCustomerdomainLeader{power}

	power.AdminCustomerdomainCustomer = &AdminCustomerdomainCustomer{power}

	power.AdminMarketMedia = &AdminMarketMedia{power}

	power.AdminMarketStore = &AdminMarketStore{power}

	power.AdminBusinessOpportunity = &AdminBusinessOpportunity{power}

	power.AdminProductPricebook = &AdminProductPricebook{power}

	power.AdminProduct = &AdminProduct{power}

	power.AdminProductCategory = &AdminProductCategory{power}

	power.AdminProductProductspecific = &AdminProductProductspecific{power}

	power.AdminProductSku = &AdminProductSku{power}

	power.AdminProductPricebookentry = &AdminProductPricebookentry{power}

	power.AdminProductArtisan = &AdminProductArtisan{power}

	power.AdminTradeAddressShipping = &AdminTradeAddressShipping{power}

	power.AdminTradeAddressBilling = &AdminTradeAddressBilling{power}

	power.AdminTradeAddressDelivery = &AdminTradeAddressDelivery{power}

	power.AdminTradeWarehouse = &AdminTradeWarehouse{power}

	power.AdminContractway = &AdminContractway{power}

	power.AdminScrmCustomer = &AdminScrmCustomer{power}

	power.AdminScrmContact = &AdminScrmContact{power}

	power.MpDictionary = &MpDictionary{power}

	power.MpCustomerAuth = &MpCustomerAuth{power}

	power.MpMarketStore = &MpMarketStore{power}

	power.MpMarketMedia = &MpMarketMedia{power}

	power.MpProductArtisan = &MpProductArtisan{power}

	power.MpProduct = &MpProduct{power}

	power.MpProduct = &MpProduct{power}

	power.MpTradeCart = &MpTradeCart{power}

	power.MpTradeOrder = &MpTradeOrder{power}

	power.AdminTradeOrder = &AdminTradeOrder{power}

	power.AdminTradePayment = &AdminTradePayment{power}

	power.MpTradeAddressShipping = &MpTradeAddressShipping{power}

	power.MpTradeAddressDelivery = &MpTradeAddressDelivery{power}

	power.MpTradeAddressBilling = &MpTradeAddressBilling{power}

	power.MpTradePayment = &MpTradePayment{power}

	power.MpDictionary = &MpDictionary{power}

	power.WebCustomerAuth = &WebCustomerAuth{power}

	power.WebCustomerAuthOa = &WebCustomerAuthOa{power}

	return power
}
