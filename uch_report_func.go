// Auto-generated code - no manual modifications
package ucbcore

// Запрос кредитного отчета
type XsRequestTypeFunc func(this *XsRequestType, ctx *ParseContext) error

var XsRequestTypeFuncInit XsRequestTypeFunc = (*XsRequestType).Init

var XsRequestTypeFuncLoad XsRequestTypeFunc = (*XsRequestType).Load

var XsRequestTypeFuncValidate XsRequestTypeFunc = (*XsRequestType).Validate

func (this *XsRequestType) Init(ctx *ParseContext) error {
  if this.Person != nil {
    if err := XsPersonSubjectTypeFuncInit(this.Person, ctx); err != nil { return err }
  }
  if this.Org != nil {
    if err := XsOrgSubjectTypeFuncInit(this.Org, ctx); err != nil { return err }
  }
  if this.Consent != nil {
    if err := XsConsentTypeFuncInit(this.Consent, ctx); err != nil { return err }
  }
  if err := XsRequestInfoTypeFuncInit(&this.RequestInfo, ctx); err != nil { return err }
  return nil
}

func (this *XsRequestType) Load(ctx *ParseContext) error {
  if this.Person != nil {
    if err := XsPersonSubjectTypeFuncLoad(this.Person, ctx); err != nil { return err }
  }
  if this.Org != nil {
    if err := XsOrgSubjectTypeFuncLoad(this.Org, ctx); err != nil { return err }
  }
  if this.Consent != nil {
    if err := XsConsentTypeFuncLoad(this.Consent, ctx); err != nil { return err }
  }
  if err := XsRequestInfoTypeFuncLoad(&this.RequestInfo, ctx); err != nil { return err }
  return nil
}

func (this *XsRequestType) Validate(ctx *ParseContext) error {
  if this.Person != nil {
    if err := XsPersonSubjectTypeFuncValidate(this.Person, ctx); err != nil { return err }
  }
  if this.Org != nil {
    if err := XsOrgSubjectTypeFuncValidate(this.Org, ctx); err != nil { return err }
  }
  if this.Consent != nil {
    if err := XsConsentTypeFuncValidate(this.Consent, ctx); err != nil { return err }
  }
  if err := XsRequestInfoTypeFuncValidate(&this.RequestInfo, ctx); err != nil { return err }
  return nil
}

// Пользователь - физическое лицо (индивидуальный предприниматель)
type XsPersonRequesterTypeFunc func(this *XsPersonRequesterType, ctx *ParseContext) error

var XsPersonRequesterTypeFuncInit XsPersonRequesterTypeFunc = (*XsPersonRequesterType).Init

var XsPersonRequesterTypeFuncLoad XsPersonRequesterTypeFunc = (*XsPersonRequesterType).Load

var XsPersonRequesterTypeFuncValidate XsPersonRequesterTypeFunc = (*XsPersonRequesterType).Validate

func (this *XsPersonRequesterType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPersonRequesterType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPersonRequesterType) Validate(ctx *ParseContext) error {
  return nil
}

// Пользователь - юридическое лицо
type XsOrgRequesterTypeFunc func(this *XsOrgRequesterType, ctx *ParseContext) error

var XsOrgRequesterTypeFuncInit XsOrgRequesterTypeFunc = (*XsOrgRequesterType).Init

var XsOrgRequesterTypeFuncLoad XsOrgRequesterTypeFunc = (*XsOrgRequesterType).Load

var XsOrgRequesterTypeFuncValidate XsOrgRequesterTypeFunc = (*XsOrgRequesterType).Validate

func (this *XsOrgRequesterType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsOrgRequesterType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsOrgRequesterType) Validate(ctx *ParseContext) error {
  return nil
}

// Сведения о согласии на получение кредитного отчета субъекта
type XsConsentTypeFunc func(this *XsConsentType, ctx *ParseContext) error

var XsConsentTypeFuncInit XsConsentTypeFunc = (*XsConsentType).Init

var XsConsentTypeFuncLoad XsConsentTypeFunc = (*XsConsentType).Load

var XsConsentTypeFuncValidate XsConsentTypeFunc = (*XsConsentType).Validate

func (this *XsConsentType) Init(ctx *ParseContext) error {
  if err := XsGivenToTypeFuncInit(&this.GivenTo, ctx); err != nil { return err }
  return nil
}

func (this *XsConsentType) Load(ctx *ParseContext) error {
  if err := XsGivenToTypeFuncLoad(&this.GivenTo, ctx); err != nil { return err }
  return nil
}

func (this *XsConsentType) Validate(ctx *ParseContext) error {
  if err := XsGivenToTypeFuncValidate(&this.GivenTo, ctx); err != nil { return err }
  return nil
}

// Запрашиваемые сведения
type XsRequestInfoTypeFunc func(this *XsRequestInfoType, ctx *ParseContext) error

var XsRequestInfoTypeFuncInit XsRequestInfoTypeFunc = (*XsRequestInfoType).Init

var XsRequestInfoTypeFuncLoad XsRequestInfoTypeFunc = (*XsRequestInfoType).Load

var XsRequestInfoTypeFuncValidate XsRequestInfoTypeFunc = (*XsRequestInfoType).Validate

func (this *XsRequestInfoType) Init(ctx *ParseContext) error {
  return nil
}

func (this *XsRequestInfoType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsRequestInfoType) Validate(ctx *ParseContext) error {
  return nil
}

// Кому выдано согласие
type XsGivenToTypeFunc func(this *XsGivenToType, ctx *ParseContext) error

var XsGivenToTypeFuncInit XsGivenToTypeFunc = (*XsGivenToType).Init

var XsGivenToTypeFuncLoad XsGivenToTypeFunc = (*XsGivenToType).Load

var XsGivenToTypeFuncValidate XsGivenToTypeFunc = (*XsGivenToType).Validate

func (this *XsGivenToType) Init(ctx *ParseContext) error {
  if this.Person != nil {
    if err := XsGivenToPersonTypeFuncInit(this.Person, ctx); err != nil { return err }
  }
  if this.Org != nil {
    if err := XsGivenToOrgTypeFuncInit(this.Org, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsGivenToType) Load(ctx *ParseContext) error {
  if this.Person != nil {
    if err := XsGivenToPersonTypeFuncLoad(this.Person, ctx); err != nil { return err }
  }
  if this.Org != nil {
    if err := XsGivenToOrgTypeFuncLoad(this.Org, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsGivenToType) Validate(ctx *ParseContext) error {
  if this.Person != nil {
    if err := XsGivenToPersonTypeFuncValidate(this.Person, ctx); err != nil { return err }
  }
  if this.Org != nil {
    if err := XsGivenToOrgTypeFuncValidate(this.Org, ctx); err != nil { return err }
  }
  return nil
}

// Согласие, выданное физическому лицу
type XsGivenToPersonTypeFunc func(this *XsGivenToPersonType, ctx *ParseContext) error

var XsGivenToPersonTypeFuncInit XsGivenToPersonTypeFunc = (*XsGivenToPersonType).Init

var XsGivenToPersonTypeFuncLoad XsGivenToPersonTypeFunc = (*XsGivenToPersonType).Load

var XsGivenToPersonTypeFuncValidate XsGivenToPersonTypeFunc = (*XsGivenToPersonType).Validate

func (this *XsGivenToPersonType) Init(ctx *ParseContext) error {
  if err := XsPersonNameTypeFuncInit(&this.Name, ctx); err != nil { return err }
  for _, v := range (*this).Tax {
    if err := XsTaxTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsGivenToPersonType) Load(ctx *ParseContext) error {
  if err := XsPersonNameTypeFuncLoad(&this.Name, ctx); err != nil { return err }
  for _, v := range (*this).Tax {
    if err := XsTaxTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsGivenToPersonType) Validate(ctx *ParseContext) error {
  if err := XsPersonNameTypeFuncValidate(&this.Name, ctx); err != nil { return err }
  for _, v := range (*this).Tax {
    if err := XsTaxTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// Согласие, выданное юридическому лицу
type XsGivenToOrgTypeFunc func(this *XsGivenToOrgType, ctx *ParseContext) error

var XsGivenToOrgTypeFuncInit XsGivenToOrgTypeFunc = (*XsGivenToOrgType).Init

var XsGivenToOrgTypeFuncLoad XsGivenToOrgTypeFunc = (*XsGivenToOrgType).Load

var XsGivenToOrgTypeFuncValidate XsGivenToOrgTypeFunc = (*XsGivenToOrgType).Validate

func (this *XsGivenToOrgType) Init(ctx *ParseContext) error {
  if err := XsOrgNameTypeFuncInit(&this.Name, ctx); err != nil { return err }
  return nil
}

func (this *XsGivenToOrgType) Load(ctx *ParseContext) error {
  if err := XsOrgNameTypeFuncLoad(&this.Name, ctx); err != nil { return err }
  return nil
}

func (this *XsGivenToOrgType) Validate(ctx *ParseContext) error {
  if err := XsOrgNameTypeFuncValidate(&this.Name, ctx); err != nil { return err }
  return nil
}

// Ответ с квитанцией о запросе кредитного отчета
type XsResultTypeFunc func(this *XsResultType, ctx *ParseContext) error

var XsResultTypeFuncInit XsResultTypeFunc = (*XsResultType).Init

var XsResultTypeFuncLoad XsResultTypeFunc = (*XsResultType).Load

var XsResultTypeFuncValidate XsResultTypeFunc = (*XsResultType).Validate

func (this *XsResultType) Init(ctx *ParseContext) error {
  return nil
}

func (this *XsResultType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsResultType) Validate(ctx *ParseContext) error {
  return nil
}

// Отчет о кредитной истории
type XsReportTypeFunc func(this *XsReportType, ctx *ParseContext) error

var XsReportTypeFuncInit XsReportTypeFunc = (*XsReportType).Init

var XsReportTypeFuncLoad XsReportTypeFunc = (*XsReportType).Load

var XsReportTypeFuncValidate XsReportTypeFunc = (*XsReportType).Validate

func (this *XsReportType) Init(ctx *ParseContext) error {
  if this.Person != nil {
    if err := XsPersonSubjectTypeFuncInit(this.Person, ctx); err != nil { return err }
  }
  if this.Org != nil {
    if err := XsOrgSubjectTypeFuncInit(this.Org, ctx); err != nil { return err }
  }
  if this.PersonInfo != nil {
    if err := XsReportPersonInfoTypeFuncInit(this.PersonInfo, ctx); err != nil { return err }
  }
  if this.OrgInfo != nil {
    if err := XsReportOrgInfoTypeFuncInit(this.OrgInfo, ctx); err != nil { return err }
  }
  for _, v := range (*this).CustomerRequest {
    if err := XsCustomerRequestTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).DealInfo {
    if err := XsReportDealInfoTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Application {
    if err := XsApplicationTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Participation {
    if err := XsParticipationTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).AppReject {
    if err := XsAppRejectTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).AvgPayment {
    if err := XsAvgPaymentTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.Rating != nil {
    if err := XsRatingTypeFuncInit(this.Rating, ctx); err != nil { return err }
  }
  if this.Score != nil {
    if err := XsScoreTypeFuncInit(this.Score, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsReportType) Load(ctx *ParseContext) error {
  if this.Person != nil {
    if err := XsPersonSubjectTypeFuncLoad(this.Person, ctx); err != nil { return err }
  }
  if this.Org != nil {
    if err := XsOrgSubjectTypeFuncLoad(this.Org, ctx); err != nil { return err }
  }
  if this.PersonInfo != nil {
    if err := XsReportPersonInfoTypeFuncLoad(this.PersonInfo, ctx); err != nil { return err }
  }
  if this.OrgInfo != nil {
    if err := XsReportOrgInfoTypeFuncLoad(this.OrgInfo, ctx); err != nil { return err }
  }
  for _, v := range (*this).CustomerRequest {
    if err := XsCustomerRequestTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).DealInfo {
    if err := XsReportDealInfoTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Application {
    if err := XsApplicationTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Participation {
    if err := XsParticipationTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).AppReject {
    if err := XsAppRejectTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).AvgPayment {
    if err := XsAvgPaymentTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if this.Rating != nil {
    if err := XsRatingTypeFuncLoad(this.Rating, ctx); err != nil { return err }
  }
  if this.Score != nil {
    if err := XsScoreTypeFuncLoad(this.Score, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsReportType) Validate(ctx *ParseContext) error {
  if this.Person != nil {
    if err := XsPersonSubjectTypeFuncValidate(this.Person, ctx); err != nil { return err }
  }
  if this.Org != nil {
    if err := XsOrgSubjectTypeFuncValidate(this.Org, ctx); err != nil { return err }
  }
  if this.PersonInfo != nil {
    if err := XsReportPersonInfoTypeFuncValidate(this.PersonInfo, ctx); err != nil { return err }
  }
  if this.OrgInfo != nil {
    if err := XsReportOrgInfoTypeFuncValidate(this.OrgInfo, ctx); err != nil { return err }
  }
  for _, v := range (*this).CustomerRequest {
    if err := XsCustomerRequestTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).DealInfo {
    if err := XsReportDealInfoTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Application {
    if err := XsApplicationTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Participation {
    if err := XsParticipationTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).AppReject {
    if err := XsAppRejectTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).AvgPayment {
    if err := XsAvgPaymentTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.Rating != nil {
    if err := XsRatingTypeFuncValidate(this.Rating, ctx); err != nil { return err }
  }
  if this.Score != nil {
    if err := XsScoreTypeFuncValidate(this.Score, ctx); err != nil { return err }
  }
  return nil
}

// Информация о физическом лице
type XsReportPersonInfoTypeFunc func(this *XsReportPersonInfoType, ctx *ParseContext) error

var XsReportPersonInfoTypeFuncInit XsReportPersonInfoTypeFunc = (*XsReportPersonInfoType).Init

var XsReportPersonInfoTypeFuncLoad XsReportPersonInfoTypeFunc = (*XsReportPersonInfoType).Load

var XsReportPersonInfoTypeFuncValidate XsReportPersonInfoTypeFunc = (*XsReportPersonInfoType).Validate

func (this *XsReportPersonInfoType) Init(ctx *ParseContext) error {
  if err := XsAddressTypeFuncInit(&this.Address, ctx); err != nil { return err }
  if this.FactAddress != nil {
    if err := XsFactAddressTypeFuncInit(this.FactAddress, ctx); err != nil { return err }
  }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.SoleProprietor != nil {
    if err := XsSoleProprietorTypeFuncInit(this.SoleProprietor, ctx); err != nil { return err }
  }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Bankruptcy {
    if err := XsBankruptcyTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Fulfillment {
    if err := XsFulfillmentTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collection {
    if err := XsCollectionTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsReportPersonInfoType) Load(ctx *ParseContext) error {
  if err := XsAddressTypeFuncLoad(&this.Address, ctx); err != nil { return err }
  if this.FactAddress != nil {
    if err := XsFactAddressTypeFuncLoad(this.FactAddress, ctx); err != nil { return err }
  }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if this.SoleProprietor != nil {
    if err := XsSoleProprietorTypeFuncLoad(this.SoleProprietor, ctx); err != nil { return err }
  }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Bankruptcy {
    if err := XsBankruptcyTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Fulfillment {
    if err := XsFulfillmentTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collection {
    if err := XsCollectionTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsReportPersonInfoType) Validate(ctx *ParseContext) error {
  if err := XsAddressTypeFuncValidate(&this.Address, ctx); err != nil { return err }
  if this.FactAddress != nil {
    if err := XsFactAddressTypeFuncValidate(this.FactAddress, ctx); err != nil { return err }
  }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.SoleProprietor != nil {
    if err := XsSoleProprietorTypeFuncValidate(this.SoleProprietor, ctx); err != nil { return err }
  }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Bankruptcy {
    if err := XsBankruptcyTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Fulfillment {
    if err := XsFulfillmentTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collection {
    if err := XsCollectionTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// Информация о юридическом лице
type XsReportOrgInfoTypeFunc func(this *XsReportOrgInfoType, ctx *ParseContext) error

var XsReportOrgInfoTypeFuncInit XsReportOrgInfoTypeFunc = (*XsReportOrgInfoType).Init

var XsReportOrgInfoTypeFuncLoad XsReportOrgInfoTypeFunc = (*XsReportOrgInfoType).Load

var XsReportOrgInfoTypeFuncValidate XsReportOrgInfoTypeFunc = (*XsReportOrgInfoType).Validate

func (this *XsReportOrgInfoType) Init(ctx *ParseContext) error {
  for _, v := range (*this).Fulfillment {
    if err := XsOrgFulfillmentTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.PrevOrg != nil {
    if err := XsPrevOrgTypeFuncInit(this.PrevOrg, ctx); err != nil { return err }
  }
  for _, v := range (*this).Bankruptcy {
    if err := XsBankruptcyTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collection {
    if err := XsCollectionTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsReportOrgInfoType) Load(ctx *ParseContext) error {
  for _, v := range (*this).Fulfillment {
    if err := XsOrgFulfillmentTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if this.PrevOrg != nil {
    if err := XsPrevOrgTypeFuncLoad(this.PrevOrg, ctx); err != nil { return err }
  }
  for _, v := range (*this).Bankruptcy {
    if err := XsBankruptcyTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collection {
    if err := XsCollectionTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsReportOrgInfoType) Validate(ctx *ParseContext) error {
  for _, v := range (*this).Fulfillment {
    if err := XsOrgFulfillmentTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.PrevOrg != nil {
    if err := XsPrevOrgTypeFuncValidate(this.PrevOrg, ctx); err != nil { return err }
  }
  for _, v := range (*this).Bankruptcy {
    if err := XsBankruptcyTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collection {
    if err := XsCollectionTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// Информация о сделке
type XsReportDealInfoTypeFunc func(this *XsReportDealInfoType, ctx *ParseContext) error

var XsReportDealInfoTypeFuncInit XsReportDealInfoTypeFunc = (*XsReportDealInfoType).Init

var XsReportDealInfoTypeFuncLoad XsReportDealInfoTypeFunc = (*XsReportDealInfoType).Load

var XsReportDealInfoTypeFuncValidate XsReportDealInfoTypeFunc = (*XsReportDealInfoType).Validate

func (this *XsReportDealInfoType) Init(ctx *ParseContext) error {
  if this.PersonMain != nil {
    if err := XsPersonDealMainTypeFuncInit(this.PersonMain, ctx); err != nil { return err }
  }
  if this.OrgMain != nil {
    if err := XsOrgDealMainTypeFuncInit(this.OrgMain, ctx); err != nil { return err }
  }
  if this.Amount != nil {
    if err := XsDealAmountTypeFuncInit(this.Amount, ctx); err != nil { return err }
  }
  if this.JointDebtors != nil {
    if err := XsJointDebtorsTypeFuncInit(this.JointDebtors, ctx); err != nil { return err }
  }
  if this.PaymentTerms != nil {
    if err := XsPaymentTermsTypeFuncInit(this.PaymentTerms, ctx); err != nil { return err }
  }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncInit(this.TotalCost, ctx); err != nil { return err }
  }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Arrear {
    if err := XsArrearTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).DueArrear {
    if err := XsDueArrearTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PastdueArrear {
    if err := XsPastdueArrearTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.MonthlyPayment != nil {
    if err := XsMonthlyPaymentTypeFuncInit(this.MonthlyPayment, ctx); err != nil { return err }
  }
  if this.SourceNonMonetaryObligation != nil {
    if err := XsSourceNonMonetaryObligationTypeFuncInit(this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  }
  if this.SubjectNonMonetaryObligation != nil {
    if err := XsSubjectNonMonetaryObligationTypeFuncInit(this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  }
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncInit(this.FundDate, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collateral {
    if err := XsCollateralTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Warranty {
    if err := XsWarrantyTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).IndepGuarantee {
    if err := XsIndepGuaranteeTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).CollateralInsurance {
    if err := XsCollateralInsuranceTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Settlement {
    if err := XsSettlementTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.Repayment != nil {
    if err := XsRepaymentTypeFuncInit(this.Repayment, ctx); err != nil { return err }
  }
  if this.Termination != nil {
    if err := XsTerminationTypeFuncInit(this.Termination, ctx); err != nil { return err }
  }
  for _, v := range (*this).Lawsuit {
    if err := XsLawsuitTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.SourceBankruptcy != nil {
    if err := XsSourceBankruptcyTypeFuncInit(this.SourceBankruptcy, ctx); err != nil { return err }
  }
  if this.SourceLiquidation != nil {
    if err := XsSourceLiquidationTypeFuncInit(this.SourceLiquidation, ctx); err != nil { return err }
  }
  for _, v := range (*this).StopInfo {
    if err := XsStopInfoTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsReportDealInfoType) Load(ctx *ParseContext) error {
  if this.PersonMain != nil {
    if err := XsPersonDealMainTypeFuncLoad(this.PersonMain, ctx); err != nil { return err }
  }
  if this.OrgMain != nil {
    if err := XsOrgDealMainTypeFuncLoad(this.OrgMain, ctx); err != nil { return err }
  }
  if this.Amount != nil {
    if err := XsDealAmountTypeFuncLoad(this.Amount, ctx); err != nil { return err }
  }
  if this.JointDebtors != nil {
    if err := XsJointDebtorsTypeFuncLoad(this.JointDebtors, ctx); err != nil { return err }
  }
  if this.PaymentTerms != nil {
    if err := XsPaymentTermsTypeFuncLoad(this.PaymentTerms, ctx); err != nil { return err }
  }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncLoad(this.TotalCost, ctx); err != nil { return err }
  }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Arrear {
    if err := XsArrearTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).DueArrear {
    if err := XsDueArrearTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PastdueArrear {
    if err := XsPastdueArrearTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if this.MonthlyPayment != nil {
    if err := XsMonthlyPaymentTypeFuncLoad(this.MonthlyPayment, ctx); err != nil { return err }
  }
  if this.SourceNonMonetaryObligation != nil {
    if err := XsSourceNonMonetaryObligationTypeFuncLoad(this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  }
  if this.SubjectNonMonetaryObligation != nil {
    if err := XsSubjectNonMonetaryObligationTypeFuncLoad(this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  }
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncLoad(this.FundDate, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collateral {
    if err := XsCollateralTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Warranty {
    if err := XsWarrantyTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).IndepGuarantee {
    if err := XsIndepGuaranteeTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).CollateralInsurance {
    if err := XsCollateralInsuranceTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Settlement {
    if err := XsSettlementTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if this.Repayment != nil {
    if err := XsRepaymentTypeFuncLoad(this.Repayment, ctx); err != nil { return err }
  }
  if this.Termination != nil {
    if err := XsTerminationTypeFuncLoad(this.Termination, ctx); err != nil { return err }
  }
  for _, v := range (*this).Lawsuit {
    if err := XsLawsuitTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if this.SourceBankruptcy != nil {
    if err := XsSourceBankruptcyTypeFuncLoad(this.SourceBankruptcy, ctx); err != nil { return err }
  }
  if this.SourceLiquidation != nil {
    if err := XsSourceLiquidationTypeFuncLoad(this.SourceLiquidation, ctx); err != nil { return err }
  }
  for _, v := range (*this).StopInfo {
    if err := XsStopInfoTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsReportDealInfoType) Validate(ctx *ParseContext) error {
  if this.PersonMain != nil {
    if err := XsPersonDealMainTypeFuncValidate(this.PersonMain, ctx); err != nil { return err }
  }
  if this.OrgMain != nil {
    if err := XsOrgDealMainTypeFuncValidate(this.OrgMain, ctx); err != nil { return err }
  }
  if this.Amount != nil {
    if err := XsDealAmountTypeFuncValidate(this.Amount, ctx); err != nil { return err }
  }
  if this.JointDebtors != nil {
    if err := XsJointDebtorsTypeFuncValidate(this.JointDebtors, ctx); err != nil { return err }
  }
  if this.PaymentTerms != nil {
    if err := XsPaymentTermsTypeFuncValidate(this.PaymentTerms, ctx); err != nil { return err }
  }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncValidate(this.TotalCost, ctx); err != nil { return err }
  }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Arrear {
    if err := XsArrearTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).DueArrear {
    if err := XsDueArrearTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PastdueArrear {
    if err := XsPastdueArrearTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.MonthlyPayment != nil {
    if err := XsMonthlyPaymentTypeFuncValidate(this.MonthlyPayment, ctx); err != nil { return err }
  }
  if this.SourceNonMonetaryObligation != nil {
    if err := XsSourceNonMonetaryObligationTypeFuncValidate(this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  }
  if this.SubjectNonMonetaryObligation != nil {
    if err := XsSubjectNonMonetaryObligationTypeFuncValidate(this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  }
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncValidate(this.FundDate, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collateral {
    if err := XsCollateralTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Warranty {
    if err := XsWarrantyTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).IndepGuarantee {
    if err := XsIndepGuaranteeTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).CollateralInsurance {
    if err := XsCollateralInsuranceTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Settlement {
    if err := XsSettlementTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.Repayment != nil {
    if err := XsRepaymentTypeFuncValidate(this.Repayment, ctx); err != nil { return err }
  }
  if this.Termination != nil {
    if err := XsTerminationTypeFuncValidate(this.Termination, ctx); err != nil { return err }
  }
  for _, v := range (*this).Lawsuit {
    if err := XsLawsuitTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.SourceBankruptcy != nil {
    if err := XsSourceBankruptcyTypeFuncValidate(this.SourceBankruptcy, ctx); err != nil { return err }
  }
  if this.SourceLiquidation != nil {
    if err := XsSourceLiquidationTypeFuncValidate(this.SourceLiquidation, ctx); err != nil { return err }
  }
  for _, v := range (*this).StopInfo {
    if err := XsStopInfoTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

