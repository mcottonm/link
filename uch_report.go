// Auto-generated code - no manual modifications
package ucbcore

// Запрос кредитного отчета
type XsRequestType struct {
  RequestId string `xml:"requestId,omitempty" key:"RequestType-requestId"`
  RequestDate string `xml:"requestDate,omitempty" key:"RequestType-requestDate"`
  CustomerUid string `xml:"customerUid,omitempty" key:"RequestType-customerUid"`
  Person *XsPersonSubjectType `xml:"person,omitempty" key:"RequestType-person"`
  Org *XsOrgSubjectType `xml:"org,omitempty" key:"RequestType-org"`
  Consent *XsConsentType `xml:"consent,omitempty" key:"RequestType-consent"`
  RequestInfo XsRequestInfoType `xml:"requestInfo,omitempty" key:"RequestType-requestInfo"`
}


// Пользователь - физическое лицо (индивидуальный предприниматель)
type XsPersonRequesterType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  FirstName string `xml:"firstName,omitempty" key:"PersonRequesterType-firstName"`
  LastName string `xml:"lastName,omitempty" key:"PersonRequesterType-lastName"`
  MidName *string `xml:"midName,omitempty" key:"PersonRequesterType-midName"`
  BirthDate string `xml:"birthDate,omitempty" key:"PersonRequesterType-birthDate"`
  BirthPlace string `xml:"birthPlace,omitempty" key:"PersonRequesterType-birthPlace"`
  IdCode string `xml:"idCode,omitempty" key:"PersonRequesterType-idCode"`
  OtherId *string `xml:"otherId,omitempty" key:"PersonRequesterType-otherId"`
  IdSeries *string `xml:"idSeries,omitempty" key:"PersonRequesterType-idSeries"`
  IdNum string `xml:"idNum,omitempty" key:"PersonRequesterType-idNum"`
  IssueDate string `xml:"issueDate,omitempty" key:"PersonRequesterType-issueDate"`
  Issuer string `xml:"issuer,omitempty" key:"PersonRequesterType-issuer"`
  DeptCode *string `xml:"deptCode,omitempty" key:"PersonRequesterType-deptCode"`
  TaxNum *string `xml:"taxNum,omitempty" key:"PersonRequesterType-taxNum"`
  TaxCode *string `xml:"taxCode,omitempty" key:"PersonRequesterType-taxCode"`
  RegNum *string `xml:"regNum,omitempty" key:"PersonRequesterType-regNum"`
  SocialId string `xml:"socialId,omitempty" key:"PersonRequesterType-socialId"`
}


// Пользователь - юридическое лицо
type XsOrgRequesterType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  RusRegInd string `xml:"rusRegInd,omitempty" key:"OrgRequesterType-rusRegInd"`
  FullName string `xml:"fullName,omitempty" key:"OrgRequesterType-fullName"`
  ShortName string `xml:"shortName,omitempty" key:"OrgRequesterType-shortName"`
  OtherName *string `xml:"otherName,omitempty" key:"OrgRequesterType-otherName"`
  Lei *string `xml:"lei,omitempty" key:"OrgRequesterType-lei"`
  RegNum *string `xml:"regNum,omitempty" key:"OrgRequesterType-regNum"`
  TaxCode *string `xml:"taxCode,omitempty" key:"OrgRequesterType-taxCode"`
  TaxNum *string `xml:"taxNum,omitempty" key:"OrgRequesterType-taxNum"`
  CustomerCode string `xml:"customerCode,omitempty" key:"OrgRequesterType-customerCode"`
}


// Сведения о согласии на получение кредитного отчета субъекта
type XsConsentType struct {
  GivenTo XsGivenToType `xml:"givenTo,omitempty" key:"ConsentType-givenTo"`
  DealDate *string `xml:"dealDate,omitempty" key:"ConsentType-dealDate"`
  HashCode *string `xml:"hashCode,omitempty" key:"ConsentType-hashCode"`
  Date string `xml:"date,omitempty" key:"ConsentType-date"`
  ExpireCode string `xml:"expireCode,omitempty" key:"ConsentType-expireCode"`
  ObtainCode string `xml:"obtainCode,omitempty" key:"ConsentType-obtainCode"`
  RequestReason []string `xml:"requestReason,omitempty" key:"ConsentType-requestReason"`
  OtherReason *string `xml:"otherReason,omitempty" key:"ConsentType-otherReason"`
  AccountabilityAcknowledged string `xml:"accountabilityAcknowledged,omitempty" key:"ConsentType-accountabilityAcknowledged"`
}


// Запрашиваемые сведения
type XsRequestInfoType struct {
  InfoCode []string `xml:"infoCode,omitempty" key:"RequestInfoType-infoCode"`
  RequestReason []string `xml:"requestReason,omitempty" key:"RequestInfoType-requestReason"`
  OtherReason *string `xml:"otherReason,omitempty" key:"RequestInfoType-otherReason"`
  Amount *string `xml:"amount,omitempty" key:"RequestInfoType-amount"`
  Currency *string `xml:"currency,omitempty" key:"RequestInfoType-currency"`
}


// Кому выдано согласие
type XsGivenToType struct {
  Person *XsGivenToPersonType `xml:"person,omitempty" key:"GivenToType-person"`
  Org *XsGivenToOrgType `xml:"org,omitempty" key:"GivenToType-org"`
}


// Согласие, выданное физическому лицу
type XsGivenToPersonType struct {
  Name XsPersonNameType `xml:"name,omitempty" key:"GivenToPersonType-name"`
  Tax []XsTaxType `xml:"tax,omitempty" key:"GivenToPersonType-tax"`
}


// Согласие, выданное юридическому лицу
type XsGivenToOrgType struct {
  Name XsOrgNameType `xml:"name,omitempty" key:"GivenToOrgType-name"`
  RegNum *string `xml:"regNum,omitempty" key:"GivenToOrgType-regNum"`
}


// Ответ с квитанцией о запросе кредитного отчета
type XsResultType struct {
  ResponseId *string `xml:"responseId,omitempty" key:"ResultType-responseId"`
  Error *string `xml:"error,omitempty" key:"ResultType-error"`
}


// Отчет о кредитной истории
type XsReportType struct {
  Person *XsPersonSubjectType `xml:"person,omitempty" key:"ReportType-person"`
  Org *XsOrgSubjectType `xml:"org,omitempty" key:"ReportType-org"`
  PersonInfo *XsReportPersonInfoType `xml:"personInfo,omitempty" key:"ReportType-personInfo"`
  OrgInfo *XsReportOrgInfoType `xml:"orgInfo,omitempty" key:"ReportType-orgInfo"`
  CustomerRequest []XsCustomerRequestType `xml:"customerRequest,omitempty" key:"ReportType-customerRequest"`
  DealInfo []XsReportDealInfoType `xml:"dealInfo,omitempty" key:"ReportType-dealInfo"`
  Application []XsApplicationType `xml:"application,omitempty" key:"ReportType-application"`
  Participation []XsParticipationType `xml:"participation,omitempty" key:"ReportType-participation"`
  AppReject []XsAppRejectType `xml:"appReject,omitempty" key:"ReportType-appReject"`
  AvgPayment []XsAvgPaymentType `xml:"avgPayment,omitempty" key:"ReportType-avgPayment"`
  Rating *XsRatingType `xml:"rating,omitempty" key:"ReportType-rating"`
  Score *XsScoreType `xml:"score,omitempty" key:"ReportType-score"`
}


// Информация о физическом лице
type XsReportPersonInfoType struct {
  Address XsAddressType `xml:"address,omitempty" key:"ReportPersonInfoType-address"`
  FactAddress *XsFactAddressType `xml:"factAddress,omitempty" key:"ReportPersonInfoType-factAddress"`
  Contact []XsContactType `xml:"contact,omitempty" key:"ReportPersonInfoType-contact"`
  SoleProprietor *XsSoleProprietorType `xml:"soleProprietor,omitempty" key:"ReportPersonInfoType-soleProprietor"`
  LegalCapacity []XsLegalCapacityType `xml:"legalCapacity,omitempty" key:"ReportPersonInfoType-legalCapacity"`
  Bankruptcy []XsBankruptcyType `xml:"bankruptcy,omitempty" key:"ReportPersonInfoType-bankruptcy"`
  Fulfillment []XsFulfillmentType `xml:"fulfillment,omitempty" key:"ReportPersonInfoType-fulfillment"`
  Collection []XsCollectionType `xml:"collection,omitempty" key:"ReportPersonInfoType-collection"`
}


// Информация о юридическом лице
type XsReportOrgInfoType struct {
  Fulfillment []XsOrgFulfillmentType `xml:"fulfillment,omitempty" key:"ReportOrgInfoType-fulfillment"`
  PrevOrg *XsPrevOrgType `xml:"prevOrg,omitempty" key:"ReportOrgInfoType-prevOrg"`
  Bankruptcy []XsBankruptcyType `xml:"bankruptcy,omitempty" key:"ReportOrgInfoType-bankruptcy"`
  Collection []XsCollectionType `xml:"collection,omitempty" key:"ReportOrgInfoType-collection"`
}


// Информация о сделке
type XsReportDealInfoType struct {
  DealUid string `xml:"dealUid,omitempty" key:"ReportDealInfoType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"ReportDealInfoType-subjectRole"`
  PersonMain *XsPersonDealMainType `xml:"personMain,omitempty" key:"ReportDealInfoType-personMain"`
  OrgMain *XsOrgDealMainType `xml:"orgMain,omitempty" key:"ReportDealInfoType-orgMain"`
  Amount *XsDealAmountType `xml:"amount,omitempty" key:"ReportDealInfoType-amount"`
  JointDebtors *XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"ReportDealInfoType-jointDebtors"`
  PaymentTerms *XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"ReportDealInfoType-paymentTerms"`
  TotalCost *XsTotalCostType `xml:"totalCost,omitempty" key:"ReportDealInfoType-totalCost"`
  Change []XsDealChangeType `xml:"change,omitempty" key:"ReportDealInfoType-change"`
  Arrear []XsArrearType `xml:"arrear,omitempty" key:"ReportDealInfoType-arrear"`
  DueArrear []XsDueArrearType `xml:"dueArrear,omitempty" key:"ReportDealInfoType-dueArrear"`
  PastdueArrear []XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"ReportDealInfoType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"ReportDealInfoType-payments"`
  MonthlyPayment *XsMonthlyPaymentType `xml:"monthlyPayment,omitempty" key:"ReportDealInfoType-monthlyPayment"`
  SourceNonMonetaryObligation *XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"ReportDealInfoType-sourceNonMonetaryObligation"`
  SubjectNonMonetaryObligation *XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"ReportDealInfoType-subjectNonMonetaryObligation"`
  FundDate *XsFundDateType `xml:"fundDate,omitempty" key:"ReportDealInfoType-fundDate"`
  Collateral []XsCollateralType `xml:"collateral,omitempty" key:"ReportDealInfoType-collateral"`
  Warranty []XsWarrantyType `xml:"warranty,omitempty" key:"ReportDealInfoType-warranty"`
  IndepGuarantee []XsIndepGuaranteeType `xml:"indepGuarantee,omitempty" key:"ReportDealInfoType-indepGuarantee"`
  CollateralInsurance []XsCollateralInsuranceType `xml:"collateralInsurance,omitempty" key:"ReportDealInfoType-collateralInsurance"`
  Settlement []XsSettlementType `xml:"settlement,omitempty" key:"ReportDealInfoType-settlement"`
  Repayment *XsRepaymentType `xml:"repayment,omitempty" key:"ReportDealInfoType-repayment"`
  Termination *XsTerminationType `xml:"termination,omitempty" key:"ReportDealInfoType-termination"`
  Lawsuit []XsLawsuitType `xml:"lawsuit,omitempty" key:"ReportDealInfoType-lawsuit"`
  SourceBankruptcy *XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"ReportDealInfoType-sourceBankruptcy"`
  SourceLiquidation *XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"ReportDealInfoType-sourceLiquidation"`
  StopInfo []XsStopInfoType `xml:"stopInfo,omitempty" key:"ReportDealInfoType-stopInfo"`
}


