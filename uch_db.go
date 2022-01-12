// Auto-generated code - no manual modifications
package ucbcore

// Доменная модель
type XsDomainModelType struct {
  Person *XsPersonSubjectType `xml:"person,omitempty" key:"DomainModelType-person"`
  Org *XsOrgSubjectType `xml:"org,omitempty" key:"DomainModelType-org"`
  SubjectInfo []XsSubjectInfoType `xml:"subjectInfo,omitempty" key:"DomainModelType-subjectInfo"`
  DealsInfo []XsDealsInfoType `xml:"dealsInfo,omitempty" key:"DomainModelType-dealsInfo"`
  CustomerInfo []XsCustomerInfoType `xml:"customerInfo,omitempty" key:"DomainModelType-customerInfo"`
}


// Доменная модель для физического лица
type XsPersonModelType struct {
  Name XsPersonNameType `xml:"name,omitempty" key:"PersonModelType-name"`
  PrevName []XsPersonPrevNameType `xml:"prevName,omitempty" key:"PersonModelType-prevName"`
  BirthInfo XsBirthInfoType `xml:"birthInfo,omitempty" key:"PersonModelType-birthInfo"`
  Id []XsPersonIdType `xml:"id,omitempty" key:"PersonModelType-id"`
  PrevId []XsPersonPrevIdType `xml:"prevId,omitempty" key:"PersonModelType-prevId"`
  Tax []XsTaxType `xml:"tax,omitempty" key:"PersonModelType-tax"`
  SocialId *XsSocialIdType `xml:"socialId,omitempty" key:"PersonModelType-socialId"`
  Address XsAddressType `xml:"address,omitempty" key:"PersonModelType-address"`
  FactAddress *XsFactAddressType `xml:"factAddress,omitempty" key:"PersonModelType-factAddress"`
  Contact []XsContactType `xml:"contact,omitempty" key:"PersonModelType-contact"`
  SoleProprietor *XsSoleProprietorType `xml:"soleProprietor,omitempty" key:"PersonModelType-soleProprietor"`
  LegalCapacity []XsLegalCapacityType `xml:"legalCapacity,omitempty" key:"PersonModelType-legalCapacity"`
  Bankruptcy []XsBankruptcyType `xml:"bankruptcy,omitempty" key:"PersonModelType-bankruptcy"`
  Fulfillment []XsFulfillmentType `xml:"fulfillment,omitempty" key:"PersonModelType-fulfillment"`
  Rating []XsRatingType `xml:"rating,omitempty" key:"PersonModelType-rating"`
  Score []XsScoreType `xml:"score,omitempty" key:"PersonModelType-score"`
  AvgPayment []XsAvgPaymentType `xml:"avgPayment,omitempty" key:"PersonModelType-avgPayment"`
  Collection []XsCollectionType `xml:"collection,omitempty" key:"PersonModelType-collection"`
  OrgCustomer []XsOrgCustomerType `xml:"orgCustomer,omitempty" key:"PersonModelType-orgCustomer"`
  PersonCustomer []XsPersonCustomerType `xml:"personCustomer,omitempty" key:"PersonModelType-personCustomer"`
  Application []XsApplicationType `xml:"application,omitempty" key:"PersonModelType-application"`
  Participation []XsParticipationType `xml:"participation,omitempty" key:"PersonModelType-participation"`
  AppReject []XsAppRejectType `xml:"appReject,omitempty" key:"PersonModelType-appReject"`
}


// Доменная модель для юридического лица
type XsOrgModelType struct {
  Name XsOrgNameType `xml:"name,omitempty" key:"OrgModelType-name"`
  Address XsOrgAddressType `xml:"address,omitempty" key:"OrgModelType-address"`
  RegNum XsOrgRegNumType `xml:"regNum,omitempty" key:"OrgModelType-regNum"`
  Tax *XsOrgTaxType `xml:"tax,omitempty" key:"OrgModelType-tax"`
  OrgChange []XsOrgChangeType `xml:"orgChange,omitempty" key:"OrgModelType-orgChange"`
  Bankruptcy []XsBankruptcyType `xml:"bankruptcy,omitempty" key:"OrgModelType-bankruptcy"`
  Fulfillment []XsOrgFulfillmentType `xml:"fulfillment,omitempty" key:"OrgModelType-fulfillment"`
  PrevOrg []XsPrevOrgType `xml:"prevOrg,omitempty" key:"OrgModelType-prevOrg"`
  Score []XsScoreType `xml:"score,omitempty" key:"OrgModelType-score"`
  Collection []XsCollectionType `xml:"collection,omitempty" key:"OrgModelType-collection"`
  OrgCustomer []XsOrgCustomerType `xml:"orgCustomer,omitempty" key:"OrgModelType-orgCustomer"`
  PersonCustomer []XsPersonCustomerType `xml:"personCustomer,omitempty" key:"OrgModelType-personCustomer"`
  Application []XsApplicationType `xml:"application,omitempty" key:"OrgModelType-application"`
  Participation []XsParticipationType `xml:"participation,omitempty" key:"OrgModelType-participation"`
  AppReject []XsAppRejectType `xml:"appReject,omitempty" key:"OrgModelType-appReject"`
}


// Доменная модель для сделки физического лица
type XsPersonDealModelType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"PersonDealModelType-dealUid"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"PersonDealModelType-main"`
  Amount *XsDealAmountType `xml:"amount,omitempty" key:"PersonDealModelType-amount"`
  JointDebtors *XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"PersonDealModelType-jointDebtors"`
  PaymentTerms *XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"PersonDealModelType-paymentTerms"`
  TotalCost *XsTotalCostType `xml:"totalCost,omitempty" key:"PersonDealModelType-totalCost"`
  Change []XsDealChangeType `xml:"change,omitempty" key:"PersonDealModelType-change"`
  FundDate *XsFundDateType `xml:"fundDate,omitempty" key:"PersonDealModelType-fundDate"`
  Arrear []XsArrearType `xml:"arrear,omitempty" key:"PersonDealModelType-arrear"`
  DueArrear []XsDueArrearType `xml:"dueArrear,omitempty" key:"PersonDealModelType-dueArrear"`
  PastdueArrear []XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"PersonDealModelType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"PersonDealModelType-payments"`
  MonthlyPayment *XsMonthlyPaymentType `xml:"monthlyPayment,omitempty" key:"PersonDealModelType-monthlyPayment"`
  SourceNonMonetaryObligation *XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"PersonDealModelType-sourceNonMonetaryObligation"`
  SubjectNonMonetaryObligation *XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"PersonDealModelType-subjectNonMonetaryObligation"`
  Collateral []XsCollateralType `xml:"collateral,omitempty" key:"PersonDealModelType-collateral"`
  Warranty []XsWarrantyType `xml:"warranty,omitempty" key:"PersonDealModelType-warranty"`
  IndepGuarantee []XsIndepGuaranteeType `xml:"indepGuarantee,omitempty" key:"PersonDealModelType-indepGuarantee"`
  CollateralInsurance []XsCollateralInsuranceType `xml:"collateralInsurance,omitempty" key:"PersonDealModelType-collateralInsurance"`
  Settlement []XsSettlementType `xml:"settlement,omitempty" key:"PersonDealModelType-settlement"`
  Repayment *XsRepaymentType `xml:"repayment,omitempty" key:"PersonDealModelType-repayment"`
  Termination *XsTerminationType `xml:"termination,omitempty" key:"PersonDealModelType-termination"`
  Lawsuit []XsLawsuitType `xml:"lawsuit,omitempty" key:"PersonDealModelType-lawsuit"`
  SourceBankruptcy *XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"PersonDealModelType-sourceBankruptcy"`
  SourceLiquidation *XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"PersonDealModelType-sourceLiquidation"`
  StopInfo []XsStopInfoType `xml:"stopInfo,omitempty" key:"PersonDealModelType-stopInfo"`
  OrgAcquirerInfo *XsOrgAcquirerInfoType `xml:"orgAcquirerInfo,omitempty" key:"PersonDealModelType-orgAcquirerInfo"`
  PersonAcquirerInfo *XsPersonAcquirerInfoType `xml:"personAcquirerInfo,omitempty" key:"PersonDealModelType-personAcquirerInfo"`
  ServiceOrg *XsServiceOrgType `xml:"serviceOrg,omitempty" key:"PersonDealModelType-serviceOrg"`
  Accounting *XsAccountingType `xml:"accounting,omitempty" key:"PersonDealModelType-accounting"`
}


// Доменная модель для сделки юридического лица
type XsOrgDealModelType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgDealModelType-dealUid"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgDealModelType-main"`
  Amount *XsDealAmountType `xml:"amount,omitempty" key:"OrgDealModelType-amount"`
  JointDebtors *XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"OrgDealModelType-jointDebtors"`
  PaymentTerms *XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"OrgDealModelType-paymentTerms"`
  Change []XsDealChangeType `xml:"change,omitempty" key:"OrgDealModelType-change"`
  FundDate *XsFundDateType `xml:"fundDate,omitempty" key:"OrgDealModelType-fundDate"`
  Arrear []XsArrearType `xml:"arrear,omitempty" key:"OrgDealModelType-arrear"`
  DueArrear []XsDueArrearType `xml:"dueArrear,omitempty" key:"OrgDealModelType-dueArrear"`
  PastdueArrear []XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"OrgDealModelType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"OrgDealModelType-payments"`
  MonthlyPayment *XsMonthlyPaymentType `xml:"monthlyPayment,omitempty" key:"OrgDealModelType-monthlyPayment"`
  SourceNonMonetaryObligation *XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"OrgDealModelType-sourceNonMonetaryObligation"`
  SubjectNonMonetaryObligation *XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"OrgDealModelType-subjectNonMonetaryObligation"`
  Collateral []XsCollateralType `xml:"collateral,omitempty" key:"OrgDealModelType-collateral"`
  Warranty []XsWarrantyType `xml:"warranty,omitempty" key:"OrgDealModelType-warranty"`
  IndepGuarantee []XsIndepGuaranteeType `xml:"indepGuarantee,omitempty" key:"OrgDealModelType-indepGuarantee"`
  CollateralInsurance []XsCollateralInsuranceType `xml:"collateralInsurance,omitempty" key:"OrgDealModelType-collateralInsurance"`
  Settlement []XsSettlementType `xml:"settlement,omitempty" key:"OrgDealModelType-settlement"`
  Repayment *XsRepaymentType `xml:"repayment,omitempty" key:"OrgDealModelType-repayment"`
  Termination *XsTerminationType `xml:"termination,omitempty" key:"OrgDealModelType-termination"`
  Lawsuit []XsLawsuitType `xml:"lawsuit,omitempty" key:"OrgDealModelType-lawsuit"`
  SourceBankruptcy *XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"OrgDealModelType-sourceBankruptcy"`
  SourceLiquidation *XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"OrgDealModelType-sourceLiquidation"`
  StopInfo []XsStopInfoType `xml:"stopInfo,omitempty" key:"OrgDealModelType-stopInfo"`
  OrgAcquirerInfo *XsOrgAcquirerInfoType `xml:"orgAcquirerInfo,omitempty" key:"OrgDealModelType-orgAcquirerInfo"`
  PersonAcquirerInfo *XsPersonAcquirerInfoType `xml:"personAcquirerInfo,omitempty" key:"OrgDealModelType-personAcquirerInfo"`
  ServiceOrg *XsServiceOrgType `xml:"serviceOrg,omitempty" key:"OrgDealModelType-serviceOrg"`
  Accounting *XsAccountingType `xml:"accounting,omitempty" key:"OrgDealModelType-accounting"`
}


// Базовый тип для информации с указанием источника
type XsSourcedInfoType struct {
  OrgSource *XsOrgSourceType `xml:"orgSource,omitempty" key:"SourcedInfoType-orgSource"`
  PersonSource *XsPersonSourceType `xml:"personSource,omitempty" key:"SourcedInfoType-personSource"`
  CommissionerSource *XsCommissionerSourceType `xml:"commissionerSource,omitempty" key:"SourcedInfoType-commissionerSource"`
}


// Информация о субъекте
type XsSubjectInfoType struct {
  OrgSource *XsOrgSourceType `xml:"orgSource,omitempty" key:"SubjectInfoType-orgSource"`
  PersonSource *XsPersonSourceType `xml:"personSource,omitempty" key:"SubjectInfoType-personSource"`
  CommissionerSource *XsCommissionerSourceType `xml:"commissionerSource,omitempty" key:"SubjectInfoType-commissionerSource"`
  Person *XsPersonModelType `xml:"person,omitempty" key:"SubjectInfoType-person"`
  Org *XsOrgModelType `xml:"org,omitempty" key:"SubjectInfoType-org"`
}


// Информация о сделках
type XsDealsInfoType struct {
  DealUid string `xml:"dealUid,omitempty" key:"DealsInfoType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"DealsInfoType-subjectRole"`
  DealInfo []XsDealInfoType `xml:"dealInfo,omitempty" key:"DealsInfoType-dealInfo"`
}


// Информация о сделке
type XsDealInfoType struct {
  OrgSource *XsOrgSourceType `xml:"orgSource,omitempty" key:"DealInfoType-orgSource"`
  PersonSource *XsPersonSourceType `xml:"personSource,omitempty" key:"DealInfoType-personSource"`
  CommissionerSource *XsCommissionerSourceType `xml:"commissionerSource,omitempty" key:"DealInfoType-commissionerSource"`
  PersonDeal *XsPersonDealModelType `xml:"personDeal,omitempty" key:"DealInfoType-personDeal"`
  OrgDeal *XsOrgDealModelType `xml:"orgDeal,omitempty" key:"DealInfoType-orgDeal"`
}


// Информация о пользователях Кредитной Истории
type XsCustomerInfoType struct {
  OrgCustomer []XsOrgCustomerType `xml:"orgCustomer,omitempty" key:"CustomerInfoType-orgCustomer"`
  PersonCustomer []XsPersonCustomerType `xml:"personCustomer,omitempty" key:"CustomerInfoType-personCustomer"`
}


