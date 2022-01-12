// Auto-generated code - no manual modifications
package ucbcore

// Пакет данных
type XsPackageType struct {
  ExternalId string `xml:"externalId,omitempty" key:"PackageType-externalId"`
  PrevExternalId *string `xml:"prevExternalId,omitempty" key:"PackageType-prevExternalId"`
  CreateDate string `xml:"createDate,omitempty" key:"PackageType-createDate"`
  TotalRecords string `xml:"totalRecords,omitempty" key:"PackageType-totalRecords"`
  TotalSubjects string `xml:"totalSubjects,omitempty" key:"PackageType-totalSubjects"`
  Source XsSourceType `xml:"source,omitempty" key:"PackageType-source"`
  Record []XsRecordType `xml:"record,omitempty" key:"PackageType-record"`
}


// Запись в пакете
type XsRecordType struct {
  OrderNum string `xml:"orderNum,omitempty" key:"RecordType-orderNum"`
  Operation string `xml:"operation,omitempty" key:"RecordType-operation"`
  OperationReason *string `xml:"operationReason,omitempty" key:"RecordType-operationReason"`
  OperationComment *string `xml:"operationComment,omitempty" key:"RecordType-operationComment"`
  PersonEvent *XsPersonEventType `xml:"personEvent,omitempty" key:"RecordType-personEvent"`
  OrgEvent *XsOrgEventType `xml:"orgEvent,omitempty" key:"RecordType-orgEvent"`
  InternalEvent *XsInternalEventType `xml:"internalEvent,omitempty" key:"RecordType-internalEvent"`
}


// Событие Кредитной Истории Физического Лица
type XsPersonEventType struct {
  Subject XsPersonSubjectType `xml:"subject,omitempty" key:"PersonEventType-subject"`
  SubjectApplied *XsSubjectAppliedType `xml:"subjectApplied,omitempty" key:"PersonEventType-subjectApplied"`
  ApplicationChanged *XsApplicationChangedType `xml:"applicationChanged,omitempty" key:"PersonEventType-applicationChanged"`
  ApplicationRejected *XsApplicationRejectedType `xml:"applicationRejected,omitempty" key:"PersonEventType-applicationRejected"`
  MonetaryDealMade *XsMonetaryDealMadeType `xml:"monetaryDealMade,omitempty" key:"PersonEventType-monetaryDealMade"`
  SourceNonMonetaryDealMade *XsSourceNonMonetaryDealMadeType `xml:"sourceNonMonetaryDealMade,omitempty" key:"PersonEventType-sourceNonMonetaryDealMade"`
  NonMonetaryDealMade *XsNonMonetaryDealMadeType `xml:"nonMonetaryDealMade,omitempty" key:"PersonEventType-nonMonetaryDealMade"`
  LeasingDealMade *XsLeasingDealMadeType `xml:"leasingDealMade,omitempty" key:"PersonEventType-leasingDealMade"`
  DebtClaimSubmitted *XsDebtClaimSubmittedType `xml:"debtClaimSubmitted,omitempty" key:"PersonEventType-debtClaimSubmitted"`
  DebtClaimChanged *XsDebtClaimChangedType `xml:"debtClaimChanged,omitempty" key:"PersonEventType-debtClaimChanged"`
  SubjectMainChanged *XsSubjectMainChangedType `xml:"subjectMainChanged,omitempty" key:"PersonEventType-subjectMainChanged"`
  SubjectSpecialChanged *XsSubjectSpecialChangedType `xml:"subjectSpecialChanged,omitempty" key:"PersonEventType-subjectSpecialChanged"`
  SubjectCapacityChanged *XsSubjectCapacityChangedType `xml:"subjectCapacityChanged,omitempty" key:"PersonEventType-subjectCapacityChanged"`
  SubjectBankruptcyChanged *XsSubjectBankruptcyChangedType `xml:"subjectBankruptcyChanged,omitempty" key:"PersonEventType-subjectBankruptcyChanged"`
  MonetaryDealChanged *XsMonetaryDealChangedType `xml:"monetaryDealChanged,omitempty" key:"PersonEventType-monetaryDealChanged"`
  NonMonetaryDealChanged *XsNonMonetaryDealChangedType `xml:"nonMonetaryDealChanged,omitempty" key:"PersonEventType-nonMonetaryDealChanged"`
  MonetaryDealPrincipalChanged *XsMonetaryDealPrincipalChangedType `xml:"monetaryDealPrincipalChanged,omitempty" key:"PersonEventType-monetaryDealPrincipalChanged"`
  NonMonetaryDealPrincipalChanged *XsNonMonetaryDealPrincipalChangedType `xml:"nonMonetaryDealPrincipalChanged,omitempty" key:"PersonEventType-nonMonetaryDealPrincipalChanged"`
  MonetaryDealPerformanceChanged *XsMonetaryDealPerformanceChangedType `xml:"monetaryDealPerformanceChanged,omitempty" key:"PersonEventType-monetaryDealPerformanceChanged"`
  NonMonetaryDealPerformanceChanged *XsNonMonetaryDealPerformanceChangedType `xml:"nonMonetaryDealPerformanceChanged,omitempty" key:"PersonEventType-nonMonetaryDealPerformanceChanged"`
  DealSecuringChanged *XsDealSecuringChangedType `xml:"dealSecuringChanged,omitempty" key:"PersonEventType-dealSecuringChanged"`
  MonetaryDealEnded *XsMonetaryDealEndedType `xml:"monetaryDealEnded,omitempty" key:"PersonEventType-monetaryDealEnded"`
  NonMonetaryDealEnded *XsNonMonetaryDealEndedType `xml:"nonMonetaryDealEnded,omitempty" key:"PersonEventType-nonMonetaryDealEnded"`
  DealClaimChanged *XsDealClaimChangedType `xml:"dealClaimChanged,omitempty" key:"PersonEventType-dealClaimChanged"`
  SourceBankruptcyStarted *XsSourceBankruptcyStartedType `xml:"sourceBankruptcyStarted,omitempty" key:"PersonEventType-sourceBankruptcyStarted"`
  SourceBankruptcyChanged *XsSourceBankruptcyChangedType `xml:"sourceBankruptcyChanged,omitempty" key:"PersonEventType-sourceBankruptcyChanged"`
  SourceBankruptcyEnded *XsSourceBankruptcyEndedType `xml:"sourceBankruptcyEnded,omitempty" key:"PersonEventType-sourceBankruptcyEnded"`
  SourceLiquidationStarted *XsSourceLiquidationStartedType `xml:"sourceLiquidationStarted,omitempty" key:"PersonEventType-sourceLiquidationStarted"`
  SourceLiquidationChanged *XsSourceLiquidationChangedType `xml:"sourceLiquidationChanged,omitempty" key:"PersonEventType-sourceLiquidationChanged"`
  SourceLiquidationEnded *XsSourceLiquidationEndedType `xml:"sourceLiquidationEnded,omitempty" key:"PersonEventType-sourceLiquidationEnded"`
  DealInfoStopped *XsDealInfoStoppedType `xml:"dealInfoStopped,omitempty" key:"PersonEventType-dealInfoStopped"`
  CreditorChanged *XsCreditorChangedType `xml:"creditorChanged,omitempty" key:"PersonEventType-creditorChanged"`
  ServiceOrgChanged *XsServiceOrgChangedType `xml:"serviceOrgChanged,omitempty" key:"PersonEventType-serviceOrgChanged"`
}


// Событие Кредитной Истории Юридического Лица
type XsOrgEventType struct {
  Subject XsOrgSubjectType `xml:"subject,omitempty" key:"OrgEventType-subject"`
  SubjectApplied *XsOrgSubjectAppliedType `xml:"subjectApplied,omitempty" key:"OrgEventType-subjectApplied"`
  ApplicationChanged *XsOrgApplicationChangedType `xml:"applicationChanged,omitempty" key:"OrgEventType-applicationChanged"`
  ApplicationRejected *XsOrgApplicationRejectedType `xml:"applicationRejected,omitempty" key:"OrgEventType-applicationRejected"`
  MonetaryDealMade *XsOrgMonetaryDealMadeType `xml:"monetaryDealMade,omitempty" key:"OrgEventType-monetaryDealMade"`
  SourceNonMonetaryDealMade *XsOrgSourceNonMonetaryDealMadeType `xml:"sourceNonMonetaryDealMade,omitempty" key:"OrgEventType-sourceNonMonetaryDealMade"`
  NonMonetaryDealMade *XsOrgNonMonetaryDealMadeType `xml:"nonMonetaryDealMade,omitempty" key:"OrgEventType-nonMonetaryDealMade"`
  LeasingDealMade *XsOrgLeasingDealMadeType `xml:"leasingDealMade,omitempty" key:"OrgEventType-leasingDealMade"`
  DebtClaimSubmitted *XsOrgDebtClaimSubmittedType `xml:"debtClaimSubmitted,omitempty" key:"OrgEventType-debtClaimSubmitted"`
  DebtClaimChanged *XsOrgDebtClaimChangedType `xml:"debtClaimChanged,omitempty" key:"OrgEventType-debtClaimChanged"`
  SubjectMainChanged *XsOrgSubjectMainChangedType `xml:"subjectMainChanged,omitempty" key:"OrgEventType-subjectMainChanged"`
  SubjectSpecialChanged *XsOrgSubjectSpecialChangedType `xml:"subjectSpecialChanged,omitempty" key:"OrgEventType-subjectSpecialChanged"`
  SubjectBankruptcyChanged *XsOrgSubjectBankruptcyChangedType `xml:"subjectBankruptcyChanged,omitempty" key:"OrgEventType-subjectBankruptcyChanged"`
  MonetaryDealChanged *XsOrgMonetaryDealChangedType `xml:"monetaryDealChanged,omitempty" key:"OrgEventType-monetaryDealChanged"`
  NonMonetaryDealChanged *XsOrgNonMonetaryDealChangedType `xml:"nonMonetaryDealChanged,omitempty" key:"OrgEventType-nonMonetaryDealChanged"`
  MonetaryDealPrincipalChanged *XsOrgMonetaryDealPrincipalChangedType `xml:"monetaryDealPrincipalChanged,omitempty" key:"OrgEventType-monetaryDealPrincipalChanged"`
  NonMonetaryDealPrincipalChanged *XsOrgNonMonetaryDealPrincipalChangedType `xml:"nonMonetaryDealPrincipalChanged,omitempty" key:"OrgEventType-nonMonetaryDealPrincipalChanged"`
  MonetaryDealPerformanceChanged *XsOrgMonetaryDealParticipationChangedType `xml:"monetaryDealPerformanceChanged,omitempty" key:"OrgEventType-monetaryDealPerformanceChanged"`
  NonMonetaryDealPerformanceChanged *XsOrgNonMonetaryDealParticipationChangedType `xml:"nonMonetaryDealPerformanceChanged,omitempty" key:"OrgEventType-nonMonetaryDealPerformanceChanged"`
  DealSecuringChanged *XsOrgDealSecuringChangedType `xml:"dealSecuringChanged,omitempty" key:"OrgEventType-dealSecuringChanged"`
  MonetaryDealEnded *XsOrgMonetaryDealEndedType `xml:"monetaryDealEnded,omitempty" key:"OrgEventType-monetaryDealEnded"`
  NonMonetaryDealEnded *XsOrgNonMonetaryDealEndedType `xml:"nonMonetaryDealEnded,omitempty" key:"OrgEventType-nonMonetaryDealEnded"`
  DealClaimChanged *XsOrgDealClaimChangedType `xml:"dealClaimChanged,omitempty" key:"OrgEventType-dealClaimChanged"`
  SourceBankruptcyStarted *XsOrgSourceBankruptcyStartedType `xml:"sourceBankruptcyStarted,omitempty" key:"OrgEventType-sourceBankruptcyStarted"`
  SourceBankruptcyChanged *XsOrgSourceBankruptcyChangedType `xml:"sourceBankruptcyChanged,omitempty" key:"OrgEventType-sourceBankruptcyChanged"`
  SourceBankruptcyEnded *XsOrgSourceBankruptcyEndedType `xml:"sourceBankruptcyEnded,omitempty" key:"OrgEventType-sourceBankruptcyEnded"`
  SourceLiquidationStarted *XsOrgSourceLiquidationStartedType `xml:"sourceLiquidationStarted,omitempty" key:"OrgEventType-sourceLiquidationStarted"`
  SourceLiquidationChanged *XsOrgSourceLiquidationChangedType `xml:"sourceLiquidationChanged,omitempty" key:"OrgEventType-sourceLiquidationChanged"`
  SourceLiquidationEnded *XsOrgSourceLiquidationEndedType `xml:"sourceLiquidationEnded,omitempty" key:"OrgEventType-sourceLiquidationEnded"`
  DealInfoStopped *XsOrgDealInfoStoppedType `xml:"dealInfoStopped,omitempty" key:"OrgEventType-dealInfoStopped"`
  CreditorChanged *XsOrgCreditorChangedType `xml:"creditorChanged,omitempty" key:"OrgEventType-creditorChanged"`
  ServiceOrgChanged *XsOrgServiceOrgChangedType `xml:"serviceOrgChanged,omitempty" key:"OrgEventType-serviceOrgChanged"`
}


// Сведения об источнике формирования кредитной истории. infoDate и один из блоков 46, 47, 48 (orgSource, personSource, commissionerSource)
type XsSourceType struct {
  InfoDate string `xml:"infoDate,omitempty" key:"SourceType-infoDate"`
  OrgSource *XsOrgSourceType `xml:"orgSource,omitempty" key:"SourceType-orgSource"`
  PersonSource *XsPersonSourceType `xml:"personSource,omitempty" key:"SourceType-personSource"`
  CommissionerSource *XsCommissionerSourceType `xml:"commissionerSource,omitempty" key:"SourceType-commissionerSource"`
}


// 1.1 Субъект обратился к источнику с предложением совершить сделку
type XsSubjectAppliedType struct {
  Application XsApplicationType `xml:"application,omitempty" key:"SubjectAppliedType-application"`
}


// 1.2 Источник одобрил обращение субъекта (направил ему оферту) или изменились сведения об обращении
type XsApplicationChangedType struct {
  Application XsApplicationType `xml:"application,omitempty" key:"ApplicationChangedType-application"`
}


// 1.3 Источник отказался от совершения сделки по обращению субъекта
type XsApplicationRejectedType struct {
  Application XsApplicationType `xml:"application,omitempty" key:"ApplicationRejectedType-application"`
  AppReject XsAppRejectType `xml:"appReject,omitempty" key:"ApplicationRejectedType-appReject"`
}


// 1.4 Субъект и источник совершили сделку, кроме договора лизинга и поручительства по лизингу для денежного обязательства субъекта
type XsMonetaryDealMadeType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"MonetaryDealMadeType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"MonetaryDealMadeType-subjectRole"`
  Address XsAddressType `xml:"address,omitempty" key:"MonetaryDealMadeType-address"`
  FactAddress XsFactAddressType `xml:"factAddress,omitempty" key:"MonetaryDealMadeType-factAddress"`
  Contact []XsContactType `xml:"contact,omitempty" key:"MonetaryDealMadeType-contact"`
  SoleProprietor XsSoleProprietorType `xml:"soleProprietor,omitempty" key:"MonetaryDealMadeType-soleProprietor"`
  LegalCapacity []XsLegalCapacityType `xml:"legalCapacity,omitempty" key:"MonetaryDealMadeType-legalCapacity"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"MonetaryDealMadeType-main"`
  JointDebtors XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"MonetaryDealMadeType-jointDebtors"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"MonetaryDealMadeType-accounting"`
  Application *XsApplicationType `xml:"application,omitempty" key:"MonetaryDealMadeType-application"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"MonetaryDealMadeType-amount"`
  PaymentTerms XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"MonetaryDealMadeType-paymentTerms"`
  TotalCost *XsTotalCostType `xml:"totalCost,omitempty" key:"MonetaryDealMadeType-totalCost"`
}


// 1.4 Субъект и источник совершили сделку, кроме договора лизинга и поручительства по лизингу для неденежного обязательства источника
type XsSourceNonMonetaryDealMadeType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"SourceNonMonetaryDealMadeType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"SourceNonMonetaryDealMadeType-subjectRole"`
  Address XsAddressType `xml:"address,omitempty" key:"SourceNonMonetaryDealMadeType-address"`
  FactAddress XsFactAddressType `xml:"factAddress,omitempty" key:"SourceNonMonetaryDealMadeType-factAddress"`
  Contact []XsContactType `xml:"contact,omitempty" key:"SourceNonMonetaryDealMadeType-contact"`
  SoleProprietor XsSoleProprietorType `xml:"soleProprietor,omitempty" key:"SourceNonMonetaryDealMadeType-soleProprietor"`
  LegalCapacity []XsLegalCapacityType `xml:"legalCapacity,omitempty" key:"SourceNonMonetaryDealMadeType-legalCapacity"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"SourceNonMonetaryDealMadeType-main"`
  JointDebtors XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"SourceNonMonetaryDealMadeType-jointDebtors"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"SourceNonMonetaryDealMadeType-accounting"`
  Application *XsApplicationType `xml:"application,omitempty" key:"SourceNonMonetaryDealMadeType-application"`
  SourceNonMonetaryObligation XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"SourceNonMonetaryDealMadeType-sourceNonMonetaryObligation"`
}


// 1.4 Субъект и источник совершили сделку, кроме договора лизинга и поручительства по лизингу для неденежного обязательства субъекта  
type XsNonMonetaryDealMadeType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"NonMonetaryDealMadeType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"NonMonetaryDealMadeType-subjectRole"`
  Address XsAddressType `xml:"address,omitempty" key:"NonMonetaryDealMadeType-address"`
  FactAddress XsFactAddressType `xml:"factAddress,omitempty" key:"NonMonetaryDealMadeType-factAddress"`
  Contact []XsContactType `xml:"contact,omitempty" key:"NonMonetaryDealMadeType-contact"`
  SoleProprietor XsSoleProprietorType `xml:"soleProprietor,omitempty" key:"NonMonetaryDealMadeType-soleProprietor"`
  LegalCapacity []XsLegalCapacityType `xml:"legalCapacity,omitempty" key:"NonMonetaryDealMadeType-legalCapacity"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"NonMonetaryDealMadeType-main"`
  JointDebtors XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"NonMonetaryDealMadeType-jointDebtors"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"NonMonetaryDealMadeType-accounting"`
  Application *XsApplicationType `xml:"application,omitempty" key:"NonMonetaryDealMadeType-application"`
  SubjectNonMonetaryObligation XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"NonMonetaryDealMadeType-subjectNonMonetaryObligation"`
}


// 1.4.1 Субъект и источник заключили договор лизинга либо поручительства по лизингу и предмет лизинга передан лизингополучателю
type XsLeasingDealMadeType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"LeasingDealMadeType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"LeasingDealMadeType-subjectRole"`
  Address XsAddressType `xml:"address,omitempty" key:"LeasingDealMadeType-address"`
  FactAddress XsFactAddressType `xml:"factAddress,omitempty" key:"LeasingDealMadeType-factAddress"`
  Contact []XsContactType `xml:"contact,omitempty" key:"LeasingDealMadeType-contact"`
  SoleProprietor XsSoleProprietorType `xml:"soleProprietor,omitempty" key:"LeasingDealMadeType-soleProprietor"`
  LegalCapacity []XsLegalCapacityType `xml:"legalCapacity,omitempty" key:"LeasingDealMadeType-legalCapacity"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"LeasingDealMadeType-main"`
  JointDebtors XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"LeasingDealMadeType-jointDebtors"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"LeasingDealMadeType-accounting"`
  Application *XsApplicationType `xml:"application,omitempty" key:"LeasingDealMadeType-application"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"LeasingDealMadeType-amount"`
  PaymentTerms XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"LeasingDealMadeType-paymentTerms"`
  FundDate XsFundDateType `xml:"fundDate,omitempty" key:"LeasingDealMadeType-fundDate"`
  Arrear XsArrearType `xml:"arrear,omitempty" key:"LeasingDealMadeType-arrear"`
  DueArrear XsDueArrearType `xml:"dueArrear,omitempty" key:"LeasingDealMadeType-dueArrear"`
  PastdueArrear XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"LeasingDealMadeType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"LeasingDealMadeType-payments"`
  SourceNonMonetaryObligation XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"LeasingDealMadeType-sourceNonMonetaryObligation"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"LeasingDealMadeType-participation"`
}


// 1.5 Для принудительного исполнения передано требование о взыскании долга по алиментам, платы за жилое помещение, коммунальные услуги или услуги связи
type XsDebtClaimSubmittedType struct {
  Address XsAddressType `xml:"address,omitempty" key:"DebtClaimSubmittedType-address"`
  FactAddress XsFactAddressType `xml:"factAddress,omitempty" key:"DebtClaimSubmittedType-factAddress"`
  Contact []XsContactType `xml:"contact,omitempty" key:"DebtClaimSubmittedType-contact"`
  SoleProprietor XsSoleProprietorType `xml:"soleProprietor,omitempty" key:"DebtClaimSubmittedType-soleProprietor"`
  LegalCapacity []XsLegalCapacityType `xml:"legalCapacity,omitempty" key:"DebtClaimSubmittedType-legalCapacity"`
  Collection XsCollectionType `xml:"collection,omitempty" key:"DebtClaimSubmittedType-collection"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"DebtClaimSubmittedType-participation"`
}


// 1.6 Изменились сведения о взыскании долга по алиментам, платы за жилое помещение, коммунальные услуги или услуги связи либо указанное требование погашено
type XsDebtClaimChangedType struct {
  Collection XsCollectionType `xml:"collection,omitempty" key:"DebtClaimChangedType-collection"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"DebtClaimChangedType-participation"`
}


// 1.7 и 1.8 Изменились сведения титульной части кредитной истории субъекта
type XsSubjectMainChangedType struct {
  Name XsPersonNameType `xml:"name,omitempty" key:"SubjectMainChangedType-name"`
  PrevName []XsPersonPrevNameType `xml:"prevName,omitempty" key:"SubjectMainChangedType-prevName"`
  BirthInfo XsBirthInfoType `xml:"birthInfo,omitempty" key:"SubjectMainChangedType-birthInfo"`
  Id []XsPersonIdType `xml:"id,omitempty" key:"SubjectMainChangedType-id"`
  PrevId []XsPersonPrevIdType `xml:"prevId,omitempty" key:"SubjectMainChangedType-prevId"`
  Tax []XsTaxType `xml:"tax,omitempty" key:"SubjectMainChangedType-tax"`
  SocialId *XsSocialIdType `xml:"socialId,omitempty" key:"SubjectMainChangedType-socialId"`
}


// 1.9 Изменились сведения о субъекте в основной части кредитной истории, кроме сведений о дееспособности, банкротстве, индивидуальном рейтинге и кредитной оценке
type XsSubjectSpecialChangedType struct {
  Address XsAddressType `xml:"address,omitempty" key:"SubjectSpecialChangedType-address"`
  FactAddress XsFactAddressType `xml:"factAddress,omitempty" key:"SubjectSpecialChangedType-factAddress"`
  Contact []XsContactType `xml:"contact,omitempty" key:"SubjectSpecialChangedType-contact"`
  SoleProprietor XsSoleProprietorType `xml:"soleProprietor,omitempty" key:"SubjectSpecialChangedType-soleProprietor"`
}


// 1.10 и 1.11 Изменились сведения о дееспособности субъекта
type XsSubjectCapacityChangedType struct {
  LegalCapacity []XsLegalCapacityType `xml:"legalCapacity,omitempty" key:"SubjectCapacityChangedType-legalCapacity"`
}


// 1.12 Изменились сведения по делу о банкротстве субъекта
type XsSubjectBankruptcyChangedType struct {
  Bankruptcy XsBankruptcyType `xml:"bankruptcy,omitempty" key:"SubjectBankruptcyChangedType-bankruptcy"`
  Fulfillment XsFulfillmentType `xml:"fulfillment,omitempty" key:"SubjectBankruptcyChangedType-fulfillment"`
}


// 1.13 Рассчитан индивидуальный рейтинг субъекта (вследствие обращения за его рейтингом или кредитным отчетом)
type XsRatingCalculatedType struct {
  Rating XsRatingType `xml:"Rating,omitempty" key:"RatingCalculatedType-Rating"`
}


// 1.14 Пользователь обратился за кредитной оценкой (скорингом) субъекта
type XsScoringRequestedType struct {
  Score XsScoreType `xml:"Score,omitempty" key:"ScoringRequestedType-Score"`
}


// 1.15 Пользователь запросил кредитный отчет субъекта
type XsReportRequestedType struct {
  CustomerRequest XsCustomerRequestType `xml:"customerRequest,omitempty" key:"ReportRequestedType-customerRequest"`
}


// 2.1 Изменились сведения об условиях обязательства субъекта для денежного обязательства
type XsMonetaryDealChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"MonetaryDealChangedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"MonetaryDealChangedType-subjectRole"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"MonetaryDealChangedType-main"`
  JointDebtors XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"MonetaryDealChangedType-jointDebtors"`
  Change []XsDealChangeType `xml:"change,omitempty" key:"MonetaryDealChangedType-change"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"MonetaryDealChangedType-accounting"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"MonetaryDealChangedType-amount"`
  PaymentTerms XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"MonetaryDealChangedType-paymentTerms"`
  FundDate XsFundDateType `xml:"fundDate,omitempty" key:"MonetaryDealChangedType-fundDate"`
  Arrear XsArrearType `xml:"arrear,omitempty" key:"MonetaryDealChangedType-arrear"`
  DueArrear XsDueArrearType `xml:"dueArrear,omitempty" key:"MonetaryDealChangedType-dueArrear"`
  PastdueArrear XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"MonetaryDealChangedType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"MonetaryDealChangedType-payments"`
  TotalCost *XsTotalCostType `xml:"totalCost,omitempty" key:"MonetaryDealChangedType-totalCost"`
  MonthlyPayment *XsMonthlyPaymentType `xml:"monthlyPayment,omitempty" key:"MonetaryDealChangedType-monthlyPayment"`
}


// 2.1 Изменились сведения об условиях обязательства субъекта для неденежного обязательства
type XsNonMonetaryDealChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"NonMonetaryDealChangedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"NonMonetaryDealChangedType-subjectRole"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"NonMonetaryDealChangedType-main"`
  JointDebtors XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"NonMonetaryDealChangedType-jointDebtors"`
  Change []XsDealChangeType `xml:"change,omitempty" key:"NonMonetaryDealChangedType-change"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"NonMonetaryDealChangedType-accounting"`
  SubjectNonMonetaryObligation XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"NonMonetaryDealChangedType-subjectNonMonetaryObligation"`
}


// 2.2 Субъекту передана сумма займа (кредита) либо субъект стал принципалом по гарантии или поручителем по сделке, кроме договора лизинга для денежного обязательства
type XsMonetaryDealPrincipalChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"MonetaryDealPrincipalChangedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"MonetaryDealPrincipalChangedType-subjectRole"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"MonetaryDealPrincipalChangedType-main"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"MonetaryDealPrincipalChangedType-participation"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"MonetaryDealPrincipalChangedType-amount"`
  PaymentTerms XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"MonetaryDealPrincipalChangedType-paymentTerms"`
  FundDate XsFundDateType `xml:"fundDate,omitempty" key:"MonetaryDealPrincipalChangedType-fundDate"`
  Arrear XsArrearType `xml:"arrear,omitempty" key:"MonetaryDealPrincipalChangedType-arrear"`
  DueArrear XsDueArrearType `xml:"dueArrear,omitempty" key:"MonetaryDealPrincipalChangedType-dueArrear"`
  PastdueArrear XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"MonetaryDealPrincipalChangedType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"MonetaryDealPrincipalChangedType-payments"`
  TotalCost *XsTotalCostType `xml:"totalCost,omitempty" key:"MonetaryDealPrincipalChangedType-totalCost"`
  MonthlyPayment *XsMonthlyPaymentType `xml:"monthlyPayment,omitempty" key:"MonetaryDealPrincipalChangedType-monthlyPayment"`
}


// 2.2 Субъекту передана сумма займа (кредита) либо субъект стал принципалом по гарантии или поручителем по сделке, кроме договора лизинга для неденежного обязательства
type XsNonMonetaryDealPrincipalChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"NonMonetaryDealPrincipalChangedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"NonMonetaryDealPrincipalChangedType-subjectRole"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"NonMonetaryDealPrincipalChangedType-main"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"NonMonetaryDealPrincipalChangedType-participation"`
  SourceNonMonetaryObligation XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"NonMonetaryDealPrincipalChangedType-sourceNonMonetaryObligation"`
}


// 2.3 Изменились сведения об исполнении обязательства субъектом, наступила ответственность поручителя или обязательство принципала возместить выплаченную сумму для денежного обязательства
type XsMonetaryDealPerformanceChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"MonetaryDealPerformanceChangedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"MonetaryDealPerformanceChangedType-subjectRole"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"MonetaryDealPerformanceChangedType-main"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"MonetaryDealPerformanceChangedType-accounting"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"MonetaryDealPerformanceChangedType-participation"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"MonetaryDealPerformanceChangedType-amount"`
  PaymentTerms XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"MonetaryDealPerformanceChangedType-paymentTerms"`
  FundDate XsFundDateType `xml:"fundDate,omitempty" key:"MonetaryDealPerformanceChangedType-fundDate"`
  Arrear XsArrearType `xml:"arrear,omitempty" key:"MonetaryDealPerformanceChangedType-arrear"`
  DueArrear XsDueArrearType `xml:"dueArrear,omitempty" key:"MonetaryDealPerformanceChangedType-dueArrear"`
  PastdueArrear XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"MonetaryDealPerformanceChangedType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"MonetaryDealPerformanceChangedType-payments"`
  TotalCost *XsTotalCostType `xml:"totalCost,omitempty" key:"MonetaryDealPerformanceChangedType-totalCost"`
  MonthlyPayment *XsMonthlyPaymentType `xml:"monthlyPayment,omitempty" key:"MonetaryDealPerformanceChangedType-monthlyPayment"`
}


// 2.3 Изменились сведения об исполнении обязательства субъектом, наступила ответственность поручителя или обязательство принципала возместить выплаченную сумму для неденежного обязательства
type XsNonMonetaryDealPerformanceChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"NonMonetaryDealPerformanceChangedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"NonMonetaryDealPerformanceChangedType-subjectRole"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"NonMonetaryDealPerformanceChangedType-main"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"NonMonetaryDealPerformanceChangedType-accounting"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"NonMonetaryDealPerformanceChangedType-participation"`
  SubjectNonMonetaryObligation XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"NonMonetaryDealPerformanceChangedType-subjectNonMonetaryObligation"`
}


// 2.4 Изменились сведения об обеспечении исполнения обязательства
type XsDealSecuringChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"DealSecuringChangedType-dealUid"`
  Collateral []XsCollateralType `xml:"collateral,omitempty" key:"DealSecuringChangedType-collateral"`
  Warranty []XsWarrantyType `xml:"warranty,omitempty" key:"DealSecuringChangedType-warranty"`
  IndepGuarantee []XsIndepGuaranteeType `xml:"indepGuarantee,omitempty" key:"DealSecuringChangedType-indepGuarantee"`
  CollateralInsurance []XsCollateralInsuranceType `xml:"collateralInsurance,omitempty" key:"DealSecuringChangedType-collateralInsurance"`
  Settlement []XsSettlementType `xml:"settlement,omitempty" key:"DealSecuringChangedType-settlement"`
  Repayment *XsRepaymentType `xml:"repayment,omitempty" key:"DealSecuringChangedType-repayment"`
}


// 2.5 Обязательство субъекта прекратилось для денежного обязательства
type XsMonetaryDealEndedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"MonetaryDealEndedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"MonetaryDealEndedType-subjectRole"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"MonetaryDealEndedType-main"`
  Termination XsTerminationType `xml:"termination,omitempty" key:"MonetaryDealEndedType-termination"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"MonetaryDealEndedType-participation"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"MonetaryDealEndedType-amount"`
  PaymentTerms XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"MonetaryDealEndedType-paymentTerms"`
  FundDate XsFundDateType `xml:"fundDate,omitempty" key:"MonetaryDealEndedType-fundDate"`
  Arrear XsArrearType `xml:"arrear,omitempty" key:"MonetaryDealEndedType-arrear"`
  DueArrear XsDueArrearType `xml:"dueArrear,omitempty" key:"MonetaryDealEndedType-dueArrear"`
  PastdueArrear XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"MonetaryDealEndedType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"MonetaryDealEndedType-payments"`
  TotalCost *XsTotalCostType `xml:"totalCost,omitempty" key:"MonetaryDealEndedType-totalCost"`
  MonthlyPayment *XsMonthlyPaymentType `xml:"monthlyPayment,omitempty" key:"MonetaryDealEndedType-monthlyPayment"`
}


// 2.5 Обязательство субъекта прекратилось для неденежного обязательства
type XsNonMonetaryDealEndedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"NonMonetaryDealEndedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"NonMonetaryDealEndedType-subjectRole"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"NonMonetaryDealEndedType-main"`
  Termination XsTerminationType `xml:"termination,omitempty" key:"NonMonetaryDealEndedType-termination"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"NonMonetaryDealEndedType-participation"`
  SubjectNonMonetaryObligation XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"NonMonetaryDealEndedType-subjectNonMonetaryObligation"`
}


// 2.6 Изменились сведения о судебном споре или требовании по обязательству
type XsDealClaimChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"DealClaimChangedType-dealUid"`
  Lawsuit []XsLawsuitType `xml:"lawsuit,omitempty" key:"DealClaimChangedType-lawsuit"`
}


// 2.7 Квалифицированное бюро получило от бюро данные для формирования сведений о среднемесячных платежах субъекта
type XsAvgPaymentsSuppliedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"AvgPaymentsSuppliedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"AvgPaymentsSuppliedType-subjectRole"`
  AvgPayments XsAvgPaymentType `xml:"avgPayments,omitempty" key:"AvgPaymentsSuppliedType-avgPayments"`
}


// 2.8.1 Конкурсное производство в отношении источника открылось
type XsSourceBankruptcyStartedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"SourceBankruptcyStartedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"SourceBankruptcyStartedType-subjectRole"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"SourceBankruptcyStartedType-main"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"SourceBankruptcyStartedType-amount"`
  SourceBankruptcy XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"SourceBankruptcyStartedType-sourceBankruptcy"`
}


// 2.8.2 В ходе конкурсного производства изменились сведения об исполнении субъектом своего обязательства или его части либо обязательство прекратилось
type XsSourceBankruptcyChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"SourceBankruptcyChangedType-dealUid"`
  SourceBankruptcy XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"SourceBankruptcyChangedType-sourceBankruptcy"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"SourceBankruptcyChangedType-participation"`
}


// 2.8.3 Конкурсное производство в отношении источника завершилось
type XsSourceBankruptcyEndedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"SourceBankruptcyEndedType-dealUid"`
  SourceBankruptcy XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"SourceBankruptcyEndedType-sourceBankruptcy"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"SourceBankruptcyEndedType-participation"`
}


// 2.9.1 Процесс ликвидации источника начался
type XsSourceLiquidationStartedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"SourceLiquidationStartedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"SourceLiquidationStartedType-subjectRole"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"SourceLiquidationStartedType-main"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"SourceLiquidationStartedType-amount"`
  SourceLiquidation XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"SourceLiquidationStartedType-sourceLiquidation"`
}


// 2.9.2 В ходе процесса ликвидации изменились сведения об исполнении субъектом своего обязательства или его части либо обязательство прекратилось
type XsSourceLiquidationChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"SourceLiquidationChangedType-dealUid"`
  SourceLiquidation XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"SourceLiquidationChangedType-sourceLiquidation"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"SourceLiquidationChangedType-participation"`
}


// 2.9.3 Процесс ликвидации источника завершился
type XsSourceLiquidationEndedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"SourceLiquidationEndedType-dealUid"`
  SourceLiquidation XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"SourceLiquidationEndedType-sourceLiquidation"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"SourceLiquidationEndedType-participation"`
}


// 2.10 Источник прекратил передачу информации по обязательству
type XsDealInfoStoppedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"DealInfoStoppedType-dealUid"`
  StopInfo XsStopInfoType `xml:"stopInfo,omitempty" key:"DealInfoStoppedType-stopInfo"`
}


// 2.11 Права кредитора по обязательству перешли к другому лицу
type XsCreditorChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"CreditorChangedType-dealUid"`
  StopInfo XsStopInfoType `xml:"stopInfo,omitempty" key:"CreditorChangedType-stopInfo"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"CreditorChangedType-accounting"`
  OrgAcquirerInfo *XsOrgAcquirerInfoType `xml:"orgAcquirerInfo,omitempty" key:"CreditorChangedType-orgAcquirerInfo"`
  PersonAcquirerInfo *XsPersonAcquirerInfoType `xml:"personAcquirerInfo,omitempty" key:"CreditorChangedType-personAcquirerInfo"`
}


// 2.12 Изменились сведения об обслуживающей организации (в частности, заключен, изменен или расторгнут договор обслуживания)
type XsServiceOrgChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"ServiceOrgChangedType-dealUid"`
  ServiceOrg XsServiceOrgType `xml:"serviceOrg,omitempty" key:"ServiceOrgChangedType-serviceOrg"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"ServiceOrgChangedType-accounting"`
}


// 1.1 Субъект обратился к источнику с предложением совершить сделку
type XsOrgSubjectAppliedType struct {
  Application XsApplicationType `xml:"application,omitempty" key:"OrgSubjectAppliedType-application"`
}


// 1.2 Источник одобрил обращение субъекта (направил ему оферту) или изменились сведения об обращении
type XsOrgApplicationChangedType struct {
  Application XsApplicationType `xml:"application,omitempty" key:"OrgApplicationChangedType-application"`
}


// 1.3 Источник отказался от совершения сделки по обращению субъекта
type XsOrgApplicationRejectedType struct {
  Application XsApplicationType `xml:"application,omitempty" key:"OrgApplicationRejectedType-application"`
  AppReject XsAppRejectType `xml:"appReject,omitempty" key:"OrgApplicationRejectedType-appReject"`
}


// 1.4 Субъект и источник совершили сделку, кроме договора лизинга и поручительства по лизингу для денежного обязательства субъекта
type XsOrgMonetaryDealMadeType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgMonetaryDealMadeType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgMonetaryDealMadeType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgMonetaryDealMadeType-main"`
  PrevOrg []XsPrevOrgType `xml:"prevOrg,omitempty" key:"OrgMonetaryDealMadeType-prevOrg"`
  JointDebtors XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"OrgMonetaryDealMadeType-jointDebtors"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"OrgMonetaryDealMadeType-accounting"`
  Application *XsApplicationType `xml:"application,omitempty" key:"OrgMonetaryDealMadeType-application"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"OrgMonetaryDealMadeType-amount"`
  PaymentTerms XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"OrgMonetaryDealMadeType-paymentTerms"`
}


// 1.4 Субъект и источник совершили сделку, кроме договора лизинга и поручительства по лизингу для неденежного обязательства источника
type XsOrgSourceNonMonetaryDealMadeType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgSourceNonMonetaryDealMadeType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgSourceNonMonetaryDealMadeType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgSourceNonMonetaryDealMadeType-main"`
  PrevOrg []XsPrevOrgType `xml:"prevOrg,omitempty" key:"OrgSourceNonMonetaryDealMadeType-prevOrg"`
  JointDebtors XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"OrgSourceNonMonetaryDealMadeType-jointDebtors"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"OrgSourceNonMonetaryDealMadeType-accounting"`
  Application *XsApplicationType `xml:"application,omitempty" key:"OrgSourceNonMonetaryDealMadeType-application"`
  SourceNonMonetaryObligation XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"OrgSourceNonMonetaryDealMadeType-sourceNonMonetaryObligation"`
}


// 1.4 Субъект и источник совершили сделку, кроме договора лизинга и поручительства по лизингу для неденежного обязательства субъекта  
type XsOrgNonMonetaryDealMadeType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgNonMonetaryDealMadeType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgNonMonetaryDealMadeType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgNonMonetaryDealMadeType-main"`
  PrevOrg []XsPrevOrgType `xml:"prevOrg,omitempty" key:"OrgNonMonetaryDealMadeType-prevOrg"`
  JointDebtors XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"OrgNonMonetaryDealMadeType-jointDebtors"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"OrgNonMonetaryDealMadeType-accounting"`
  Application *XsApplicationType `xml:"application,omitempty" key:"OrgNonMonetaryDealMadeType-application"`
  SubjectNonMonetaryObligation XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"OrgNonMonetaryDealMadeType-subjectNonMonetaryObligation"`
}


// 1.4.1 Субъект и источник заключили договор лизинга либо поручительства по лизингу и предмет лизинга передан лизингополучателю
type XsOrgLeasingDealMadeType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgLeasingDealMadeType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgLeasingDealMadeType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgLeasingDealMadeType-main"`
  PrevOrg []XsPrevOrgType `xml:"prevOrg,omitempty" key:"OrgLeasingDealMadeType-prevOrg"`
  JointDebtors XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"OrgLeasingDealMadeType-jointDebtors"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"OrgLeasingDealMadeType-accounting"`
  Application *XsApplicationType `xml:"application,omitempty" key:"OrgLeasingDealMadeType-application"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"OrgLeasingDealMadeType-amount"`
  PaymentTerms XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"OrgLeasingDealMadeType-paymentTerms"`
  FundDate XsFundDateType `xml:"fundDate,omitempty" key:"OrgLeasingDealMadeType-fundDate"`
  Arrear XsArrearType `xml:"arrear,omitempty" key:"OrgLeasingDealMadeType-arrear"`
  DueArrear XsDueArrearType `xml:"dueArrear,omitempty" key:"OrgLeasingDealMadeType-dueArrear"`
  PastdueArrear XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"OrgLeasingDealMadeType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"OrgLeasingDealMadeType-payments"`
  SourceNonMonetaryObligation XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"OrgLeasingDealMadeType-sourceNonMonetaryObligation"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgLeasingDealMadeType-participation"`
}


// 1.5 Для принудительного исполнения передано требование о взыскании долга по алиментам, платы за жилое помещение, коммунальные услуги или услуги связи
type XsOrgDebtClaimSubmittedType struct {
  PrevOrg []XsPrevOrgType `xml:"prevOrg,omitempty" key:"OrgDebtClaimSubmittedType-prevOrg"`
  Collection XsCollectionType `xml:"collection,omitempty" key:"OrgDebtClaimSubmittedType-collection"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgDebtClaimSubmittedType-participation"`
}


// 1.6 Изменились сведения о взыскании долга по алиментам, платы за жилое помещение, коммунальные услуги или услуги связи либо указанное требование погашено
type XsOrgDebtClaimChangedType struct {
  Collection XsCollectionType `xml:"collection,omitempty" key:"OrgDebtClaimChangedType-collection"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgDebtClaimChangedType-participation"`
}


// 1.7 и 1.8 Изменились сведения титульной части кредитной истории субъекта
type XsOrgSubjectMainChangedType struct {
  Name XsOrgNameType `xml:"name,omitempty" key:"OrgSubjectMainChangedType-name"`
  Address XsOrgAddressType `xml:"address,omitempty" key:"OrgSubjectMainChangedType-address"`
  RegNum XsOrgRegNumType `xml:"regNum,omitempty" key:"OrgSubjectMainChangedType-regNum"`
  Tax *XsOrgTaxType `xml:"tax,omitempty" key:"OrgSubjectMainChangedType-tax"`
  OrgChange []XsOrgChangeType `xml:"orgChange,omitempty" key:"OrgSubjectMainChangedType-orgChange"`
}


// 1.9 Изменились сведения о субъекте в основной части кредитной истории, кроме сведений о дееспособности, банкротстве, индивидуальном рейтинге и кредитной оценке
type XsOrgSubjectSpecialChangedType struct {
  PrevOrg []XsPrevOrgType `xml:"prevOrg,omitempty" key:"OrgSubjectSpecialChangedType-prevOrg"`
}


// 1.10 и 1.11 Изменились сведения о дееспособности субъекта
type XsOrgSubjectCapacityChangedType struct {
  LegalCapacity XsLegalCapacityType `xml:"legalCapacity,omitempty" key:"OrgSubjectCapacityChangedType-legalCapacity"`
}


// 1.12 Изменились сведения по делу о банкротстве субъекта
type XsOrgSubjectBankruptcyChangedType struct {
  Bankruptcy XsBankruptcyType `xml:"bankruptcy,omitempty" key:"OrgSubjectBankruptcyChangedType-bankruptcy"`
  Fulfillment XsOrgFulfillmentType `xml:"fulfillment,omitempty" key:"OrgSubjectBankruptcyChangedType-fulfillment"`
}


// 2.1 Изменились сведения об условиях обязательства субъекта для денежного обязательства
type XsOrgMonetaryDealChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgMonetaryDealChangedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgMonetaryDealChangedType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgMonetaryDealChangedType-main"`
  JointDebtors XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"OrgMonetaryDealChangedType-jointDebtors"`
  Change []XsDealChangeType `xml:"change,omitempty" key:"OrgMonetaryDealChangedType-change"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"OrgMonetaryDealChangedType-accounting"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"OrgMonetaryDealChangedType-amount"`
  PaymentTerms XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"OrgMonetaryDealChangedType-paymentTerms"`
  FundDate XsFundDateType `xml:"fundDate,omitempty" key:"OrgMonetaryDealChangedType-fundDate"`
  Arrear XsArrearType `xml:"arrear,omitempty" key:"OrgMonetaryDealChangedType-arrear"`
  DueArrear XsDueArrearType `xml:"dueArrear,omitempty" key:"OrgMonetaryDealChangedType-dueArrear"`
  PastdueArrear XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"OrgMonetaryDealChangedType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"OrgMonetaryDealChangedType-payments"`
}


// 2.1 Изменились сведения об условиях обязательства субъекта для неденежного обязательства
type XsOrgNonMonetaryDealChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgNonMonetaryDealChangedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgNonMonetaryDealChangedType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgNonMonetaryDealChangedType-main"`
  JointDebtors XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"OrgNonMonetaryDealChangedType-jointDebtors"`
  Change []XsDealChangeType `xml:"change,omitempty" key:"OrgNonMonetaryDealChangedType-change"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"OrgNonMonetaryDealChangedType-accounting"`
  SubjectNonMonetaryObligation XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"OrgNonMonetaryDealChangedType-subjectNonMonetaryObligation"`
}


// 2.2 Субъекту передана сумма займа (кредита) либо субъект стал принципалом по гарантии или поручителем по сделке, кроме договора лизинга для денежного обязательства
type XsOrgMonetaryDealPrincipalChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgMonetaryDealPrincipalChangedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgMonetaryDealPrincipalChangedType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgMonetaryDealPrincipalChangedType-main"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgMonetaryDealPrincipalChangedType-participation"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"OrgMonetaryDealPrincipalChangedType-amount"`
  PaymentTerms XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"OrgMonetaryDealPrincipalChangedType-paymentTerms"`
  FundDate XsFundDateType `xml:"fundDate,omitempty" key:"OrgMonetaryDealPrincipalChangedType-fundDate"`
  Arrear XsArrearType `xml:"arrear,omitempty" key:"OrgMonetaryDealPrincipalChangedType-arrear"`
  DueArrear XsDueArrearType `xml:"dueArrear,omitempty" key:"OrgMonetaryDealPrincipalChangedType-dueArrear"`
  PastdueArrear XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"OrgMonetaryDealPrincipalChangedType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"OrgMonetaryDealPrincipalChangedType-payments"`
}


// 2.2 Субъекту передана сумма займа (кредита) либо субъект стал принципалом по гарантии или поручителем по сделке, кроме договора лизинга для денежного обязательства
type XsOrgNonMonetaryDealPrincipalChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgNonMonetaryDealPrincipalChangedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgNonMonetaryDealPrincipalChangedType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgNonMonetaryDealPrincipalChangedType-main"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgNonMonetaryDealPrincipalChangedType-participation"`
  SourceNonMonetaryObligation XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"OrgNonMonetaryDealPrincipalChangedType-sourceNonMonetaryObligation"`
}


// 2.3 Изменились сведения об исполнении обязательства субъектом, наступила ответственность поручителя или обязательство принципала возместить выплаченную сумму для денежного обязательства
type XsOrgMonetaryDealParticipationChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgMonetaryDealParticipationChangedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgMonetaryDealParticipationChangedType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgMonetaryDealParticipationChangedType-main"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"OrgMonetaryDealParticipationChangedType-accounting"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgMonetaryDealParticipationChangedType-participation"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"OrgMonetaryDealParticipationChangedType-amount"`
  PaymentTerms XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"OrgMonetaryDealParticipationChangedType-paymentTerms"`
  FundDate XsFundDateType `xml:"fundDate,omitempty" key:"OrgMonetaryDealParticipationChangedType-fundDate"`
  Arrear XsArrearType `xml:"arrear,omitempty" key:"OrgMonetaryDealParticipationChangedType-arrear"`
  DueArrear XsDueArrearType `xml:"dueArrear,omitempty" key:"OrgMonetaryDealParticipationChangedType-dueArrear"`
  PastdueArrear XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"OrgMonetaryDealParticipationChangedType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"OrgMonetaryDealParticipationChangedType-payments"`
}


// 2.3 Изменились сведения об исполнении обязательства субъектом, наступила ответственность поручителя или обязательство принципала возместить выплаченную сумму для неденежного обязательства
type XsOrgNonMonetaryDealParticipationChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgNonMonetaryDealParticipationChangedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgNonMonetaryDealParticipationChangedType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgNonMonetaryDealParticipationChangedType-main"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"OrgNonMonetaryDealParticipationChangedType-accounting"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgNonMonetaryDealParticipationChangedType-participation"`
  SubjectNonMonetaryObligation XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"OrgNonMonetaryDealParticipationChangedType-subjectNonMonetaryObligation"`
}


// 2.4 Изменились сведения об обеспечении исполнения обязательства
type XsOrgDealSecuringChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgDealSecuringChangedType-dealUid"`
  Collateral []XsCollateralType `xml:"collateral,omitempty" key:"OrgDealSecuringChangedType-collateral"`
  Warranty []XsWarrantyType `xml:"warranty,omitempty" key:"OrgDealSecuringChangedType-warranty"`
  IndepGuarantee []XsIndepGuaranteeType `xml:"indepGuarantee,omitempty" key:"OrgDealSecuringChangedType-indepGuarantee"`
  CollateralInsurance []XsCollateralInsuranceType `xml:"collateralInsurance,omitempty" key:"OrgDealSecuringChangedType-collateralInsurance"`
  Settlement []XsSettlementType `xml:"settlement,omitempty" key:"OrgDealSecuringChangedType-settlement"`
  Repayment *XsRepaymentType `xml:"repayment,omitempty" key:"OrgDealSecuringChangedType-repayment"`
}


// 2.5 Обязательство субъекта прекратилось для денежного обязательства
type XsOrgMonetaryDealEndedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgMonetaryDealEndedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgMonetaryDealEndedType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgMonetaryDealEndedType-main"`
  Termination XsTerminationType `xml:"termination,omitempty" key:"OrgMonetaryDealEndedType-termination"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgMonetaryDealEndedType-participation"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"OrgMonetaryDealEndedType-amount"`
  PaymentTerms XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"OrgMonetaryDealEndedType-paymentTerms"`
  FundDate XsFundDateType `xml:"fundDate,omitempty" key:"OrgMonetaryDealEndedType-fundDate"`
  Arrear XsArrearType `xml:"arrear,omitempty" key:"OrgMonetaryDealEndedType-arrear"`
  DueArrear XsDueArrearType `xml:"dueArrear,omitempty" key:"OrgMonetaryDealEndedType-dueArrear"`
  PastdueArrear XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"OrgMonetaryDealEndedType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"OrgMonetaryDealEndedType-payments"`
}


// 2.5 Обязательство субъекта прекратилось для неденежного обязательства
type XsOrgNonMonetaryDealEndedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgNonMonetaryDealEndedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgNonMonetaryDealEndedType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgNonMonetaryDealEndedType-main"`
  Termination XsTerminationType `xml:"termination,omitempty" key:"OrgNonMonetaryDealEndedType-termination"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgNonMonetaryDealEndedType-participation"`
  SubjectNonMonetaryObligation XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"OrgNonMonetaryDealEndedType-subjectNonMonetaryObligation"`
}


// 2.6 Изменились сведения о судебном споре или требовании по обязательству
type XsOrgDealClaimChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgDealClaimChangedType-dealUid"`
  Lawsuit []XsLawsuitType `xml:"lawsuit,omitempty" key:"OrgDealClaimChangedType-lawsuit"`
}


// 2.8.1 Конкурсное производство в отношении источника открылось
type XsOrgSourceBankruptcyStartedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgSourceBankruptcyStartedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgSourceBankruptcyStartedType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgSourceBankruptcyStartedType-main"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"OrgSourceBankruptcyStartedType-amount"`
  SourceBankruptcy XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"OrgSourceBankruptcyStartedType-sourceBankruptcy"`
}


// 2.8.2 В ходе конкурсного производства изменились сведения об исполнении субъектом своего обязательства или его части либо обязательство прекратилось
type XsOrgSourceBankruptcyChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgSourceBankruptcyChangedType-dealUid"`
  SourceBankruptcy XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"OrgSourceBankruptcyChangedType-sourceBankruptcy"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgSourceBankruptcyChangedType-participation"`
}


// 2.8.3 Конкурсное производство в отношении источника завершилось
type XsOrgSourceBankruptcyEndedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgSourceBankruptcyEndedType-dealUid"`
  SourceBankruptcy XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"OrgSourceBankruptcyEndedType-sourceBankruptcy"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgSourceBankruptcyEndedType-participation"`
}


// 2.9.1 Процесс ликвидации источника начался
type XsOrgSourceLiquidationStartedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgSourceLiquidationStartedType-dealUid"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"OrgSourceLiquidationStartedType-subjectRole"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgSourceLiquidationStartedType-main"`
  Amount XsDealAmountType `xml:"amount,omitempty" key:"OrgSourceLiquidationStartedType-amount"`
  SourceLiquidation XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"OrgSourceLiquidationStartedType-sourceLiquidation"`
}


// 2.9.2 В ходе процесса ликвидации изменились сведения об исполнении субъектом своего обязательства или его части либо обязательство прекратилось
type XsOrgSourceLiquidationChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgSourceLiquidationChangedType-dealUid"`
  SourceLiquidation XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"OrgSourceLiquidationChangedType-sourceLiquidation"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgSourceLiquidationChangedType-participation"`
}


// 2.9.3 Процесс ликвидации источника завершился
type XsOrgSourceLiquidationEndedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgSourceLiquidationEndedType-dealUid"`
  SourceLiquidation XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"OrgSourceLiquidationEndedType-sourceLiquidation"`
  Participation XsParticipationType `xml:"participation,omitempty" key:"OrgSourceLiquidationEndedType-participation"`
}


// 2.10 Источник прекратил передачу информации по обязательству
type XsOrgDealInfoStoppedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgDealInfoStoppedType-dealUid"`
  StopInfo XsStopInfoType `xml:"stopInfo,omitempty" key:"OrgDealInfoStoppedType-stopInfo"`
}


// 2.11 Права кредитора по обязательству перешли к другому лицу
type XsOrgCreditorChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgCreditorChangedType-dealUid"`
  StopInfo XsStopInfoType `xml:"stopInfo,omitempty" key:"OrgCreditorChangedType-stopInfo"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"OrgCreditorChangedType-accounting"`
  OrgAcquirerInfo *XsOrgAcquirerInfoType `xml:"orgAcquirerInfo,omitempty" key:"OrgCreditorChangedType-orgAcquirerInfo"`
  PersonAcquirerInfo *XsPersonAcquirerInfoType `xml:"personAcquirerInfo,omitempty" key:"OrgCreditorChangedType-personAcquirerInfo"`
}


// 2.12 Изменились сведения об обслуживающей организации (в частности, заключен, изменен или расторгнут договор обслуживания)
type XsOrgServiceOrgChangedType struct {
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgServiceOrgChangedType-dealUid"`
  ServiceOrg XsServiceOrgType `xml:"serviceOrg,omitempty" key:"OrgServiceOrgChangedType-serviceOrg"`
  Accounting XsAccountingType `xml:"accounting,omitempty" key:"OrgServiceOrgChangedType-accounting"`
}


// Внутреннее событие Кредитной - используется только Кредитным Бюро
type XsInternalEventType struct {
  Subject *XsPersonSubjectType `xml:"subject,omitempty" key:"InternalEventType-subject"`
  RatingCalculated *XsRatingCalculatedType `xml:"ratingCalculated,omitempty" key:"InternalEventType-ratingCalculated"`
  ScoringRequested *XsScoringRequestedType `xml:"scoringRequested,omitempty" key:"InternalEventType-scoringRequested"`
  ReportRequested *XsReportRequestedType `xml:"reportRequested,omitempty" key:"InternalEventType-reportRequested"`
  AvgPaymentsSupplied *XsAvgPaymentsSuppliedType `xml:"avgPaymentsSupplied,omitempty" key:"InternalEventType-avgPaymentsSupplied"`
  ImportPerson *XsImportPersonType `xml:"importPerson,omitempty" key:"InternalEventType-importPerson"`
  ImportOrg *XsImportOrgType `xml:"importOrg,omitempty" key:"InternalEventType-importOrg"`
}


// Импорт КИ физического лица
type XsImportPersonType struct {
  Name XsPersonNameType `xml:"name,omitempty" key:"ImportPersonType-name"`
  PrevName []XsPersonPrevNameType `xml:"prevName,omitempty" key:"ImportPersonType-prevName"`
  BirthInfo XsBirthInfoType `xml:"birthInfo,omitempty" key:"ImportPersonType-birthInfo"`
  Id []XsPersonIdType `xml:"id,omitempty" key:"ImportPersonType-id"`
  PrevId []XsPersonPrevIdType `xml:"prevId,omitempty" key:"ImportPersonType-prevId"`
  Tax []XsTaxType `xml:"tax,omitempty" key:"ImportPersonType-tax"`
  SocialId *XsSocialIdType `xml:"socialId,omitempty" key:"ImportPersonType-socialId"`
  Address XsAddressType `xml:"address,omitempty" key:"ImportPersonType-address"`
  FactAddress *XsFactAddressType `xml:"factAddress,omitempty" key:"ImportPersonType-factAddress"`
  Contact []XsContactType `xml:"contact,omitempty" key:"ImportPersonType-contact"`
  SoleProprietor *XsSoleProprietorType `xml:"soleProprietor,omitempty" key:"ImportPersonType-soleProprietor"`
  LegalCapacity []XsLegalCapacityType `xml:"legalCapacity,omitempty" key:"ImportPersonType-legalCapacity"`
  Bankruptcy []XsBankruptcyType `xml:"bankruptcy,omitempty" key:"ImportPersonType-bankruptcy"`
  Fulfillment []XsFulfillmentType `xml:"fulfillment,omitempty" key:"ImportPersonType-fulfillment"`
  Rating []XsRatingType `xml:"rating,omitempty" key:"ImportPersonType-rating"`
  Score []XsScoreType `xml:"score,omitempty" key:"ImportPersonType-score"`
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"ImportPersonType-dealUid"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"ImportPersonType-main"`
  Amount *XsDealAmountType `xml:"amount,omitempty" key:"ImportPersonType-amount"`
  JointDebtors *XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"ImportPersonType-jointDebtors"`
  PaymentTerms *XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"ImportPersonType-paymentTerms"`
  TotalCost *XsTotalCostType `xml:"totalCost,omitempty" key:"ImportPersonType-totalCost"`
  Change []XsDealChangeType `xml:"change,omitempty" key:"ImportPersonType-change"`
  FundDate *XsFundDateType `xml:"fundDate,omitempty" key:"ImportPersonType-fundDate"`
  Arrear []XsArrearType `xml:"arrear,omitempty" key:"ImportPersonType-arrear"`
  DueArrear []XsDueArrearType `xml:"dueArrear,omitempty" key:"ImportPersonType-dueArrear"`
  PastdueArrear []XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"ImportPersonType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"ImportPersonType-payments"`
  MonthlyPayment *XsMonthlyPaymentType `xml:"monthlyPayment,omitempty" key:"ImportPersonType-monthlyPayment"`
  SourceNonMonetaryObligation *XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"ImportPersonType-sourceNonMonetaryObligation"`
  SubjectNonMonetaryObligation *XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"ImportPersonType-subjectNonMonetaryObligation"`
  Collateral []XsCollateralType `xml:"collateral,omitempty" key:"ImportPersonType-collateral"`
  Warranty []XsWarrantyType `xml:"warranty,omitempty" key:"ImportPersonType-warranty"`
  IndepGuarantee []XsIndepGuaranteeType `xml:"indepGuarantee,omitempty" key:"ImportPersonType-indepGuarantee"`
  CollateralInsurance []XsCollateralInsuranceType `xml:"collateralInsurance,omitempty" key:"ImportPersonType-collateralInsurance"`
  Settlement []XsSettlementType `xml:"settlement,omitempty" key:"ImportPersonType-settlement"`
  Repayment *XsRepaymentType `xml:"repayment,omitempty" key:"ImportPersonType-repayment"`
  Termination *XsTerminationType `xml:"termination,omitempty" key:"ImportPersonType-termination"`
  Lawsuit []XsLawsuitType `xml:"lawsuit,omitempty" key:"ImportPersonType-lawsuit"`
  AvgPayment []XsAvgPaymentType `xml:"avgPayment,omitempty" key:"ImportPersonType-avgPayment"`
  SourceBankruptcy *XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"ImportPersonType-sourceBankruptcy"`
  SourceLiquidation *XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"ImportPersonType-sourceLiquidation"`
  Collection []XsCollectionType `xml:"collection,omitempty" key:"ImportPersonType-collection"`
  CustomerRequest []XsCustomerRequestType `xml:"customerRequest,omitempty" key:"ImportPersonType-customerRequest"`
  StopInfo []XsStopInfoType `xml:"stopInfo,omitempty" key:"ImportPersonType-stopInfo"`
  InfoDate string `xml:"infoDate,omitempty" key:"ImportPersonType-infoDate"`
  OrgSource *XsOrgSourceType `xml:"orgSource,omitempty" key:"ImportPersonType-orgSource"`
  PersonSource *XsPersonSourceType `xml:"personSource,omitempty" key:"ImportPersonType-personSource"`
  CommissionerSource *XsCommissionerSourceType `xml:"commissionerSource,omitempty" key:"ImportPersonType-commissionerSource"`
  OrgCustomer []XsOrgCustomerType `xml:"orgCustomer,omitempty" key:"ImportPersonType-orgCustomer"`
  PersonCustomer []XsPersonCustomerType `xml:"personCustomer,omitempty" key:"ImportPersonType-personCustomer"`
  OrgAcquirerInfo *XsOrgAcquirerInfoType `xml:"orgAcquirerInfo,omitempty" key:"ImportPersonType-orgAcquirerInfo"`
  PersonAcquirerInfo *XsPersonAcquirerInfoType `xml:"personAcquirerInfo,omitempty" key:"ImportPersonType-personAcquirerInfo"`
  ServiceOrg *XsServiceOrgType `xml:"serviceOrg,omitempty" key:"ImportPersonType-serviceOrg"`
  Accounting *XsAccountingType `xml:"accounting,omitempty" key:"ImportPersonType-accounting"`
  Application []XsApplicationType `xml:"application,omitempty" key:"ImportPersonType-application"`
  Participation []XsParticipationType `xml:"participation,omitempty" key:"ImportPersonType-participation"`
  AppReject []XsAppRejectType `xml:"appReject,omitempty" key:"ImportPersonType-appReject"`
  SubjectRole *string `xml:"subjectRole,omitempty" key:"ImportPersonType-subjectRole"`
}


// Импорт КИ юридического лица
type XsImportOrgType struct {
  Name XsOrgNameType `xml:"name,omitempty" key:"ImportOrgType-name"`
  Address XsOrgAddressType `xml:"address,omitempty" key:"ImportOrgType-address"`
  RegNum XsOrgRegNumType `xml:"regNum,omitempty" key:"ImportOrgType-regNum"`
  Tax *XsOrgTaxType `xml:"tax,omitempty" key:"ImportOrgType-tax"`
  OrgChange []XsOrgChangeType `xml:"orgChange,omitempty" key:"ImportOrgType-orgChange"`
  Bankruptcy []XsBankruptcyType `xml:"bankruptcy,omitempty" key:"ImportOrgType-bankruptcy"`
  Fulfillment []XsOrgFulfillmentType `xml:"fulfillment,omitempty" key:"ImportOrgType-fulfillment"`
  PrevOrg []XsPrevOrgType `xml:"prevOrg,omitempty" key:"ImportOrgType-prevOrg"`
  Score []XsScoreType `xml:"score,omitempty" key:"ImportOrgType-score"`
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"ImportOrgType-dealUid"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"ImportOrgType-main"`
  Amount *XsDealAmountType `xml:"amount,omitempty" key:"ImportOrgType-amount"`
  JointDebtors *XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"ImportOrgType-jointDebtors"`
  PaymentTerms *XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"ImportOrgType-paymentTerms"`
  Change []XsDealChangeType `xml:"change,omitempty" key:"ImportOrgType-change"`
  FundDate *XsFundDateType `xml:"fundDate,omitempty" key:"ImportOrgType-fundDate"`
  Arrear []XsArrearType `xml:"arrear,omitempty" key:"ImportOrgType-arrear"`
  DueArrear []XsDueArrearType `xml:"dueArrear,omitempty" key:"ImportOrgType-dueArrear"`
  PastdueArrear []XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"ImportOrgType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"ImportOrgType-payments"`
  MonthlyPayment *XsMonthlyPaymentType `xml:"monthlyPayment,omitempty" key:"ImportOrgType-monthlyPayment"`
  SourceNonMonetaryObligation *XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"ImportOrgType-sourceNonMonetaryObligation"`
  SubjectNonMonetaryObligation *XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"ImportOrgType-subjectNonMonetaryObligation"`
  Collateral []XsCollateralType `xml:"collateral,omitempty" key:"ImportOrgType-collateral"`
  Warranty []XsWarrantyType `xml:"warranty,omitempty" key:"ImportOrgType-warranty"`
  IndepGuarantee []XsIndepGuaranteeType `xml:"indepGuarantee,omitempty" key:"ImportOrgType-indepGuarantee"`
  CollateralInsurance []XsCollateralInsuranceType `xml:"collateralInsurance,omitempty" key:"ImportOrgType-collateralInsurance"`
  Settlement []XsSettlementType `xml:"settlement,omitempty" key:"ImportOrgType-settlement"`
  Repayment *XsRepaymentType `xml:"repayment,omitempty" key:"ImportOrgType-repayment"`
  Termination *XsTerminationType `xml:"termination,omitempty" key:"ImportOrgType-termination"`
  Lawsuit []XsLawsuitType `xml:"lawsuit,omitempty" key:"ImportOrgType-lawsuit"`
  SourceBankruptcy *XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"ImportOrgType-sourceBankruptcy"`
  SourceLiquidation *XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"ImportOrgType-sourceLiquidation"`
  Collection *XsCollectionType `xml:"collection,omitempty" key:"ImportOrgType-collection"`
  InfoDate string `xml:"infoDate,omitempty" key:"ImportOrgType-infoDate"`
  OrgSource *XsOrgSourceType `xml:"orgSource,omitempty" key:"ImportOrgType-orgSource"`
  PersonSource *XsPersonSourceType `xml:"personSource,omitempty" key:"ImportOrgType-personSource"`
  CommissionerSource *XsCommissionerSourceType `xml:"commissionerSource,omitempty" key:"ImportOrgType-commissionerSource"`
  OrgCustomer []XsOrgCustomerType `xml:"orgCustomer,omitempty" key:"ImportOrgType-orgCustomer"`
  PersonCustomer []XsPersonCustomerType `xml:"personCustomer,omitempty" key:"ImportOrgType-personCustomer"`
  OrgAcquirerInfo *XsOrgAcquirerInfoType `xml:"orgAcquirerInfo,omitempty" key:"ImportOrgType-orgAcquirerInfo"`
  PersonAcquirerInfo *XsPersonAcquirerInfoType `xml:"personAcquirerInfo,omitempty" key:"ImportOrgType-personAcquirerInfo"`
  ServiceOrg *XsServiceOrgType `xml:"serviceOrg,omitempty" key:"ImportOrgType-serviceOrg"`
  Accounting *XsAccountingType `xml:"accounting,omitempty" key:"ImportOrgType-accounting"`
  Application []XsApplicationType `xml:"application,omitempty" key:"ImportOrgType-application"`
  Participation []XsParticipationType `xml:"participation,omitempty" key:"ImportOrgType-participation"`
  AppReject []XsAppRejectType `xml:"appReject,omitempty" key:"ImportOrgType-appReject"`
  SubjectRole *string `xml:"subjectRole,omitempty" key:"ImportOrgType-subjectRole"`
}


// Кредитная история физического лица
type XsPersonCreditHistoryType struct {
  Name XsPersonNameType `xml:"name,omitempty" key:"PersonCreditHistoryType-name"`
  PrevName []XsPersonPrevNameType `xml:"prevName,omitempty" key:"PersonCreditHistoryType-prevName"`
  BirthInfo XsBirthInfoType `xml:"birthInfo,omitempty" key:"PersonCreditHistoryType-birthInfo"`
  Id []XsPersonIdType `xml:"id,omitempty" key:"PersonCreditHistoryType-id"`
  PrevId []XsPersonPrevIdType `xml:"prevId,omitempty" key:"PersonCreditHistoryType-prevId"`
  Tax []XsTaxType `xml:"tax,omitempty" key:"PersonCreditHistoryType-tax"`
  SocialId *XsSocialIdType `xml:"socialId,omitempty" key:"PersonCreditHistoryType-socialId"`
  Address XsAddressType `xml:"address,omitempty" key:"PersonCreditHistoryType-address"`
  FactAddress *XsFactAddressType `xml:"factAddress,omitempty" key:"PersonCreditHistoryType-factAddress"`
  Contact []XsContactType `xml:"contact,omitempty" key:"PersonCreditHistoryType-contact"`
  SoleProprietor *XsSoleProprietorType `xml:"soleProprietor,omitempty" key:"PersonCreditHistoryType-soleProprietor"`
  LegalCapacity []XsLegalCapacityType `xml:"legalCapacity,omitempty" key:"PersonCreditHistoryType-legalCapacity"`
  Bankruptcy []XsBankruptcyType `xml:"bankruptcy,omitempty" key:"PersonCreditHistoryType-bankruptcy"`
  Fulfillment []XsFulfillmentType `xml:"fulfillment,omitempty" key:"PersonCreditHistoryType-fulfillment"`
  Rating []XsRatingType `xml:"rating,omitempty" key:"PersonCreditHistoryType-rating"`
  Score []XsScoreType `xml:"score,omitempty" key:"PersonCreditHistoryType-score"`
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"PersonCreditHistoryType-dealUid"`
  Main XsPersonDealMainType `xml:"main,omitempty" key:"PersonCreditHistoryType-main"`
  Amount *XsDealAmountType `xml:"amount,omitempty" key:"PersonCreditHistoryType-amount"`
  JointDebtors *XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"PersonCreditHistoryType-jointDebtors"`
  PaymentTerms *XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"PersonCreditHistoryType-paymentTerms"`
  TotalCost *XsTotalCostType `xml:"totalCost,omitempty" key:"PersonCreditHistoryType-totalCost"`
  Change []XsDealChangeType `xml:"change,omitempty" key:"PersonCreditHistoryType-change"`
  FundDate *XsFundDateType `xml:"fundDate,omitempty" key:"PersonCreditHistoryType-fundDate"`
  Arrear []XsArrearType `xml:"arrear,omitempty" key:"PersonCreditHistoryType-arrear"`
  DueArrear []XsDueArrearType `xml:"dueArrear,omitempty" key:"PersonCreditHistoryType-dueArrear"`
  PastdueArrear []XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"PersonCreditHistoryType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"PersonCreditHistoryType-payments"`
  MonthlyPayment *XsMonthlyPaymentType `xml:"monthlyPayment,omitempty" key:"PersonCreditHistoryType-monthlyPayment"`
  SourceNonMonetaryObligation *XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"PersonCreditHistoryType-sourceNonMonetaryObligation"`
  SubjectNonMonetaryObligation *XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"PersonCreditHistoryType-subjectNonMonetaryObligation"`
  Collateral []XsCollateralType `xml:"collateral,omitempty" key:"PersonCreditHistoryType-collateral"`
  Warranty []XsWarrantyType `xml:"warranty,omitempty" key:"PersonCreditHistoryType-warranty"`
  IndepGuarantee []XsIndepGuaranteeType `xml:"indepGuarantee,omitempty" key:"PersonCreditHistoryType-indepGuarantee"`
  CollateralInsurance []XsCollateralInsuranceType `xml:"collateralInsurance,omitempty" key:"PersonCreditHistoryType-collateralInsurance"`
  Settlement []XsSettlementType `xml:"settlement,omitempty" key:"PersonCreditHistoryType-settlement"`
  Repayment *XsRepaymentType `xml:"repayment,omitempty" key:"PersonCreditHistoryType-repayment"`
  Termination *XsTerminationType `xml:"termination,omitempty" key:"PersonCreditHistoryType-termination"`
  Lawsuit []XsLawsuitType `xml:"lawsuit,omitempty" key:"PersonCreditHistoryType-lawsuit"`
  AvgPayment []XsAvgPaymentType `xml:"avgPayment,omitempty" key:"PersonCreditHistoryType-avgPayment"`
  SourceBankruptcy *XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"PersonCreditHistoryType-sourceBankruptcy"`
  SourceLiquidation *XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"PersonCreditHistoryType-sourceLiquidation"`
  Collection []XsCollectionType `xml:"collection,omitempty" key:"PersonCreditHistoryType-collection"`
  CustomerRequest []XsCustomerRequestType `xml:"customerRequest,omitempty" key:"PersonCreditHistoryType-customerRequest"`
  StopInfo []XsStopInfoType `xml:"stopInfo,omitempty" key:"PersonCreditHistoryType-stopInfo"`
  InfoDate string `xml:"infoDate,omitempty" key:"PersonCreditHistoryType-infoDate"`
  OrgSource *XsOrgSourceType `xml:"orgSource,omitempty" key:"PersonCreditHistoryType-orgSource"`
  PersonSource *XsPersonSourceType `xml:"personSource,omitempty" key:"PersonCreditHistoryType-personSource"`
  CommissionerSource *XsCommissionerSourceType `xml:"commissionerSource,omitempty" key:"PersonCreditHistoryType-commissionerSource"`
  OrgCustomer []XsOrgCustomerType `xml:"orgCustomer,omitempty" key:"PersonCreditHistoryType-orgCustomer"`
  PersonCustomer []XsPersonCustomerType `xml:"personCustomer,omitempty" key:"PersonCreditHistoryType-personCustomer"`
  OrgAcquirerInfo *XsOrgAcquirerInfoType `xml:"orgAcquirerInfo,omitempty" key:"PersonCreditHistoryType-orgAcquirerInfo"`
  PersonAcquirerInfo *XsPersonAcquirerInfoType `xml:"personAcquirerInfo,omitempty" key:"PersonCreditHistoryType-personAcquirerInfo"`
  ServiceOrg *XsServiceOrgType `xml:"serviceOrg,omitempty" key:"PersonCreditHistoryType-serviceOrg"`
  Accounting *XsAccountingType `xml:"accounting,omitempty" key:"PersonCreditHistoryType-accounting"`
  Application []XsApplicationType `xml:"application,omitempty" key:"PersonCreditHistoryType-application"`
  Participation []XsParticipationType `xml:"participation,omitempty" key:"PersonCreditHistoryType-participation"`
  AppReject []XsAppRejectType `xml:"appReject,omitempty" key:"PersonCreditHistoryType-appReject"`
}


// Блок 1. Имя
type XsPersonNameType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  FirstName string `xml:"firstName,omitempty" key:"PersonNameType-firstName"`
  LastName string `xml:"lastName,omitempty" key:"PersonNameType-lastName"`
  MidName *string `xml:"midName,omitempty" key:"PersonNameType-midName"`
}


// Блок 2. Предыдущее имя
type XsPersonPrevNameType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  FirstName *string `xml:"firstName,omitempty" key:"PersonPrevNameType-firstName"`
  LastName *string `xml:"lastName,omitempty" key:"PersonPrevNameType-lastName"`
  MidName *string `xml:"midName,omitempty" key:"PersonPrevNameType-midName"`
  Date *string `xml:"date,omitempty" key:"PersonPrevNameType-date"`
}


// Блок 3. Дата и место рождения
type XsBirthInfoType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  BirthDate string `xml:"birthDate,omitempty" key:"BirthInfoType-birthDate"`
  CountryCode string `xml:"countryCode,omitempty" key:"BirthInfoType-countryCode"`
  BirthPlace string `xml:"birthPlace,omitempty" key:"BirthInfoType-birthPlace"`
}


// Блок 4. Документ, удостоверяющий личность; Блок 5. Документ, ранее удостоверявший личность
type XsPersonIdType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  CountryCode string `xml:"countryCode,omitempty" key:"PersonIdType-countryCode"`
  OtherCountry *string `xml:"otherCountry,omitempty" key:"PersonIdType-otherCountry"`
  IdCode string `xml:"idCode,omitempty" key:"PersonIdType-idCode"`
  OtherId *string `xml:"otherId,omitempty" key:"PersonIdType-otherId"`
  IdSeries *string `xml:"idSeries,omitempty" key:"PersonIdType-idSeries"`
  IdNum string `xml:"idNum,omitempty" key:"PersonIdType-idNum"`
  IssueDate string `xml:"issueDate,omitempty" key:"PersonIdType-issueDate"`
  Issuer string `xml:"issuer,omitempty" key:"PersonIdType-issuer"`
  DeptCode *string `xml:"deptCode,omitempty" key:"PersonIdType-deptCode"`
  EndDate *string `xml:"endDate,omitempty" key:"PersonIdType-endDate"`
}


// Блок 5. Документ, ранее удостоверявший личность
type XsPersonPrevIdType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  CountryCode *string `xml:"countryCode,omitempty" key:"PersonPrevIdType-countryCode"`
  OtherCountry *string `xml:"otherCountry,omitempty" key:"PersonPrevIdType-otherCountry"`
  IdCode *string `xml:"idCode,omitempty" key:"PersonPrevIdType-idCode"`
  OtherId *string `xml:"otherId,omitempty" key:"PersonPrevIdType-otherId"`
  IdSeries *string `xml:"idSeries,omitempty" key:"PersonPrevIdType-idSeries"`
  IdNum *string `xml:"idNum,omitempty" key:"PersonPrevIdType-idNum"`
  IssueDate *string `xml:"issueDate,omitempty" key:"PersonPrevIdType-issueDate"`
  Issuer *string `xml:"issuer,omitempty" key:"PersonPrevIdType-issuer"`
  DeptCode *string `xml:"deptCode,omitempty" key:"PersonPrevIdType-deptCode"`
  EndDate *string `xml:"endDate,omitempty" key:"PersonPrevIdType-endDate"`
}


// Блок 6. (юр Блок 6) Номер налогоплательщика и регистрационный номер
type XsTaxType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  TaxCode *string `xml:"taxCode,omitempty" key:"TaxType-taxCode"`
  TaxNum *string `xml:"taxNum,omitempty" key:"TaxType-taxNum"`
  RegNum *string `xml:"regNum,omitempty" key:"TaxType-regNum"`
}


// Блок 7. СНИЛС
type XsSocialIdType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Id string `xml:"id,omitempty" key:"SocialIdType-id"`
}


// Блок 8. Регистрация физического лица по месту жительства или пребывания
type XsAddressType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Postal *string `xml:"postal,omitempty" key:"AddressType-postal"`
  CountryCode *string `xml:"countryCode,omitempty" key:"AddressType-countryCode"`
  OtherCountry *string `xml:"otherCountry,omitempty" key:"AddressType-otherCountry"`
  AddrNum *string `xml:"addrNum,omitempty" key:"AddressType-addrNum"`
  LocationCode *string `xml:"locationCode,omitempty" key:"AddressType-locationCode"`
  OtherLocation *string `xml:"otherLocation,omitempty" key:"AddressType-otherLocation"`
  Street *string `xml:"street,omitempty" key:"AddressType-street"`
  House *string `xml:"house,omitempty" key:"AddressType-house"`
  Estate *string `xml:"estate,omitempty" key:"AddressType-estate"`
  Block *string `xml:"block,omitempty" key:"AddressType-block"`
  Building *string `xml:"building,omitempty" key:"AddressType-building"`
  Apartment *string `xml:"apartment,omitempty" key:"AddressType-apartment"`
  RegAddrCode string `xml:"regAddrCode,omitempty" key:"AddressType-regAddrCode"`
  RegDate *string `xml:"regDate,omitempty" key:"AddressType-regDate"`
  RegAuthority *string `xml:"regAuthority,omitempty" key:"AddressType-regAuthority"`
  RegAuthorityCode *string `xml:"regAuthorityCode,omitempty" key:"AddressType-regAuthorityCode"`
}


// Блок 9. Фактическое место жительства
type XsFactAddressType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Postal *string `xml:"postal,omitempty" key:"FactAddressType-postal"`
  CountryCode *string `xml:"countryCode,omitempty" key:"FactAddressType-countryCode"`
  OtherCountry *string `xml:"otherCountry,omitempty" key:"FactAddressType-otherCountry"`
  AddrNum *string `xml:"addrNum,omitempty" key:"FactAddressType-addrNum"`
  LocationCode *string `xml:"locationCode,omitempty" key:"FactAddressType-locationCode"`
  OtherLocation *string `xml:"otherLocation,omitempty" key:"FactAddressType-otherLocation"`
  Street *string `xml:"street,omitempty" key:"FactAddressType-street"`
  House *string `xml:"house,omitempty" key:"FactAddressType-house"`
  Estate *string `xml:"estate,omitempty" key:"FactAddressType-estate"`
  Block *string `xml:"block,omitempty" key:"FactAddressType-block"`
  Building *string `xml:"building,omitempty" key:"FactAddressType-building"`
  Apartment *string `xml:"apartment,omitempty" key:"FactAddressType-apartment"`
}


// Блок 10. Контактные данные
type XsContactType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Phone *string `xml:"phone,omitempty" key:"ContactType-phone"`
  PhoneComment *string `xml:"phoneComment,omitempty" key:"ContactType-phoneComment"`
  Email *string `xml:"email,omitempty" key:"ContactType-email"`
}


// Блок 11. Государственная регистрация в качестве индивидуального предпринимателя
type XsSoleProprietorType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  RegNum *string `xml:"regNum,omitempty" key:"SoleProprietorType-regNum"`
  RegDate *string `xml:"regDate,omitempty" key:"SoleProprietorType-regDate"`
}


// Блок 12. Сведения о дееспособности
type XsLegalCapacityType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Capacity string `xml:"capacity,omitempty" key:"LegalCapacityType-capacity"`
  Date *string `xml:"date,omitempty" key:"LegalCapacityType-date"`
  Num *string `xml:"num,omitempty" key:"LegalCapacityType-num"`
  Court *string `xml:"court,omitempty" key:"LegalCapacityType-court"`
}


// Блок 13. (юр Блок 6) Сведения по делу о несостоятельности (банкротстве)
type XsBankruptcyType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Stage *string `xml:"stage,omitempty" key:"BankruptcyType-stage"`
  Date *string `xml:"date,omitempty" key:"BankruptcyType-date"`
  PublishLink *string `xml:"publishLink,omitempty" key:"BankruptcyType-publishLink"`
  IllegalActInd *string `xml:"illegalActInd,omitempty" key:"BankruptcyType-illegalActInd"`
  IllegalActDate *string `xml:"illegalActDate,omitempty" key:"BankruptcyType-illegalActDate"`
  FakeActInd *string `xml:"fakeActInd,omitempty" key:"BankruptcyType-fakeActInd"`
  FakeActDate *string `xml:"fakeActDate,omitempty" key:"BankruptcyType-fakeActDate"`
  Info *string `xml:"info,omitempty" key:"BankruptcyType-info"`
}


// Блок 14. Сведения о завершении расчетов с кредиторами и освобождении субъекта от исполнения обязательств в связи с банкротством
type XsFulfillmentType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Date *string `xml:"date,omitempty" key:"FulfillmentType-date"`
  SettleInd *string `xml:"settleInd,omitempty" key:"FulfillmentType-settleInd"`
  SettleDate *string `xml:"settleDate,omitempty" key:"FulfillmentType-settleDate"`
  ResumeDate *string `xml:"resumeDate,omitempty" key:"FulfillmentType-resumeDate"`
}


// Блок 15. Индивидуальный рейтинг субъекта
type XsRatingType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Rating *string `xml:"rating,omitempty" key:"RatingType-rating"`
  Date *string `xml:"date,omitempty" key:"RatingType-date"`
  InfluenceFactor []string `xml:"influenceFactor,omitempty" key:"RatingType-influenceFactor"`
  Influence []string `xml:"influence,omitempty" key:"RatingType-influence"`
  CalcInabilityReason *string `xml:"calcInabilityReason,omitempty" key:"RatingType-calcInabilityReason"`
  OtherInabilityReasons *string `xml:"otherInabilityReasons,omitempty" key:"RatingType-otherInabilityReasons"`
}


// Блок 16. (юр Блок 9) Кредитная оценка (скоринг)
type XsScoreType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Score string `xml:"score,omitempty" key:"ScoreType-score"`
  Date string `xml:"date,omitempty" key:"ScoreType-date"`
}


// Блок 17. (юр Блок 10) Уникальный идентификатор договора (сделки)
type XsDealUidType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Uid string `xml:"uid,omitempty" key:"DealUidType-uid"`
}


// Блок 18. Общие сведения о сделке
type XsPersonDealMainType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Date string `xml:"date,omitempty" key:"PersonDealMainType-date"`
  DealType string `xml:"dealType,omitempty" key:"PersonDealMainType-dealType"`
  LoanKind *string `xml:"loanKind,omitempty" key:"PersonDealMainType-loanKind"`
  LoanPurpose []string `xml:"loanPurpose,omitempty" key:"PersonDealMainType-loanPurpose"`
  CardInd string `xml:"cardInd,omitempty" key:"PersonDealMainType-cardInd"`
  NovationInd string `xml:"novationInd,omitempty" key:"PersonDealMainType-novationInd"`
  SourceObligationInd *string `xml:"sourceObligationInd,omitempty" key:"PersonDealMainType-sourceObligationInd"`
  SubjectObligationInd *string `xml:"subjectObligationInd,omitempty" key:"PersonDealMainType-subjectObligationInd"`
  EndDate string `xml:"endDate,omitempty" key:"PersonDealMainType-endDate"`
  ConsumerLoanInd string `xml:"consumerLoanInd,omitempty" key:"PersonDealMainType-consumerLoanInd"`
}


// Блок 19. (юр Блок 12) Сумма и валюта обязательства
type XsDealAmountType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Amount string `xml:"amount,omitempty" key:"DealAmountType-amount"`
  Currency string `xml:"currency,omitempty" key:"DealAmountType-currency"`
  SecuredAmount *string `xml:"securedAmount,omitempty" key:"DealAmountType-securedAmount"`
}


// Блок 20. (юр Блок 13) Сведения о солидарных должниках
type XsJointDebtorsType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Number *string `xml:"number,omitempty" key:"JointDebtorsType-number"`
}


// Блок 21. (юр Блок 14) Сведения об условиях платежей
type XsPaymentTermsType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  PrincipalAmount string `xml:"principalAmount,omitempty" key:"PaymentTermsType-principalAmount"`
  PrincipalDate *string `xml:"principalDate,omitempty" key:"PaymentTermsType-principalDate"`
  InterestAmount *string `xml:"interestAmount,omitempty" key:"PaymentTermsType-interestAmount"`
  InterestDate *string `xml:"interestDate,omitempty" key:"PaymentTermsType-interestDate"`
  Frequency *string `xml:"frequency,omitempty" key:"PaymentTermsType-frequency"`
  MinCardAmount *string `xml:"minCardAmount,omitempty" key:"PaymentTermsType-minCardAmount"`
  GraceStartDate *string `xml:"graceStartDate,omitempty" key:"PaymentTermsType-graceStartDate"`
  GraceEndDate *string `xml:"graceEndDate,omitempty" key:"PaymentTermsType-graceEndDate"`
  InterestEndDate *string `xml:"interestEndDate,omitempty" key:"PaymentTermsType-interestEndDate"`
}


// Блок 22. Полная стоимость потребительского кредита (займа)
type XsTotalCostType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Percents string `xml:"percents,omitempty" key:"TotalCostType-percents"`
  Money string `xml:"money,omitempty" key:"TotalCostType-money"`
  CalcDate string `xml:"calcDate,omitempty" key:"TotalCostType-calcDate"`
}


// Блок 23. (юр Блок 15) Сведения об изменении договора
type XsDealChangeType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Date *string `xml:"date,omitempty" key:"DealChangeType-date"`
  ChangeType *string `xml:"changeType,omitempty" key:"DealChangeType-changeType"`
  SubType *string `xml:"subType,omitempty" key:"DealChangeType-subType"`
  OtherDescription *string `xml:"otherDescription,omitempty" key:"DealChangeType-otherDescription"`
  StartDate *string `xml:"startDate,omitempty" key:"DealChangeType-startDate"`
  EndDate *string `xml:"endDate,omitempty" key:"DealChangeType-endDate"`
  FactEndDate *string `xml:"factEndDate,omitempty" key:"DealChangeType-factEndDate"`
  EndReason *string `xml:"endReason,omitempty" key:"DealChangeType-endReason"`
  CurrencyRate *string `xml:"currencyRate,omitempty" key:"DealChangeType-currencyRate"`
}


// Блок 24. (юр Блок 16) Дата передачи финансирования субъекту или возникновения обеспечения исполнения обязательства
type XsFundDateType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Date string `xml:"date,omitempty" key:"FundDateType-date"`
}


// Блок 25. (юр Блок 17) Сведения о задолженности
type XsArrearType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  StartAmount *string `xml:"startAmount,omitempty" key:"ArrearType-startAmount"`
  LastPaymentInd *string `xml:"lastPaymentInd,omitempty" key:"ArrearType-lastPaymentInd"`
  Amount *string `xml:"amount,omitempty" key:"ArrearType-amount"`
  PrincipalAmount *string `xml:"principalAmount,omitempty" key:"ArrearType-principalAmount"`
  InterestAmount *string `xml:"interestAmount,omitempty" key:"ArrearType-interestAmount"`
  OtherAmount *string `xml:"otherAmount,omitempty" key:"ArrearType-otherAmount"`
  Date *string `xml:"date,omitempty" key:"ArrearType-date"`
  GraceInd *string `xml:"graceInd,omitempty" key:"ArrearType-graceInd"`
}


// Блок 26. (юр Блок 18) Сведения о срочной задолженности
type XsDueArrearType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  StartDate *string `xml:"startDate,omitempty" key:"DueArrearType-startDate"`
  LastPaymentInd string `xml:"lastPaymentInd,omitempty" key:"DueArrearType-lastPaymentInd"`
  Amount string `xml:"amount,omitempty" key:"DueArrearType-amount"`
  PrincipalAmount *string `xml:"principalAmount,omitempty" key:"DueArrearType-principalAmount"`
  InterestAmount *string `xml:"interestAmount,omitempty" key:"DueArrearType-interestAmount"`
  OtherAmount *string `xml:"otherAmount,omitempty" key:"DueArrearType-otherAmount"`
  Date *string `xml:"date,omitempty" key:"DueArrearType-date"`
}


// Блок 27. (юр Блок 19) Сведения о просроченной задолженности
type XsPastdueArrearType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  StartDate *string `xml:"startDate,omitempty" key:"PastdueArrearType-startDate"`
  LastPaymentInd string `xml:"lastPaymentInd,omitempty" key:"PastdueArrearType-lastPaymentInd"`
  Amount string `xml:"amount,omitempty" key:"PastdueArrearType-amount"`
  PrincipalAmount *string `xml:"principalAmount,omitempty" key:"PastdueArrearType-principalAmount"`
  InterestAmount *string `xml:"interestAmount,omitempty" key:"PastdueArrearType-interestAmount"`
  OtherAmount *string `xml:"otherAmount,omitempty" key:"PastdueArrearType-otherAmount"`
  Date *string `xml:"date,omitempty" key:"PastdueArrearType-date"`
  PrincipalMissDate *string `xml:"principalMissDate,omitempty" key:"PastdueArrearType-principalMissDate"`
  InterestMissDate *string `xml:"interestMissDate,omitempty" key:"PastdueArrearType-interestMissDate"`
}


// Блок 28. (юр Блок 20) Сведения о внесении платежей
type XsPaymentsType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Date *string `xml:"date,omitempty" key:"PaymentsType-date"`
  LastAmount string `xml:"lastAmount,omitempty" key:"PaymentsType-lastAmount"`
  LastPrincipalAmount *string `xml:"lastPrincipalAmount,omitempty" key:"PaymentsType-lastPrincipalAmount"`
  LastInterestAmount *string `xml:"lastInterestAmount,omitempty" key:"PaymentsType-lastInterestAmount"`
  LastOtherAmount *string `xml:"lastOtherAmount,omitempty" key:"PaymentsType-lastOtherAmount"`
  TotalAmount string `xml:"totalAmount,omitempty" key:"PaymentsType-totalAmount"`
  TotalPrincipalAmount string `xml:"totalPrincipalAmount,omitempty" key:"PaymentsType-totalPrincipalAmount"`
  TotalInterestAmount string `xml:"totalInterestAmount,omitempty" key:"PaymentsType-totalInterestAmount"`
  TotalOtherAmount string `xml:"totalOtherAmount,omitempty" key:"PaymentsType-totalOtherAmount"`
  AmountCompliance string `xml:"amountCompliance,omitempty" key:"PaymentsType-amountCompliance"`
  ScheduleCompliance string `xml:"scheduleCompliance,omitempty" key:"PaymentsType-scheduleCompliance"`
  PastduePeriod string `xml:"pastduePeriod,omitempty" key:"PaymentsType-pastduePeriod"`
}


// Блок 29. Величина среднемесячного платежа по договору займа (кредита) и дата ее расчета
type XsMonthlyPaymentType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Amount string `xml:"amount,omitempty" key:"MonthlyPaymentType-amount"`
  Date string `xml:"date,omitempty" key:"MonthlyPaymentType-date"`
}


// Блок 30. (юр Блок 21) Сведения о неденежном обязательстве источника
type XsSourceNonMonetaryObligationType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Asset string `xml:"asset,omitempty" key:"SourceNonMonetaryObligationType-asset"`
  AssetKind string `xml:"assetKind,omitempty" key:"SourceNonMonetaryObligationType-assetKind"`
  Object string `xml:"object,omitempty" key:"SourceNonMonetaryObligationType-object"`
  TransferDate string `xml:"transferDate,omitempty" key:"SourceNonMonetaryObligationType-transferDate"`
}


// Блок 31. (юр Блок 22) Сведения о неденежном обязательстве субъекта
type XsSubjectNonMonetaryObligationType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Asset string `xml:"asset,omitempty" key:"SubjectNonMonetaryObligationType-asset"`
  Object string `xml:"object,omitempty" key:"SubjectNonMonetaryObligationType-object"`
  Terms string `xml:"terms,omitempty" key:"SubjectNonMonetaryObligationType-terms"`
  ComplianceInd string `xml:"complianceInd,omitempty" key:"SubjectNonMonetaryObligationType-complianceInd"`
}


// Блок 32. (юр Блок 23) Сведения о залоге
type XsCollateralType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  AssetKind *string `xml:"assetKind,omitempty" key:"CollateralType-assetKind"`
  AssetId *string `xml:"assetId,omitempty" key:"CollateralType-assetId"`
  ContractDate *string `xml:"contractDate,omitempty" key:"CollateralType-contractDate"`
  AssetValue *string `xml:"assetValue,omitempty" key:"CollateralType-assetValue"`
  Currency *string `xml:"currency,omitempty" key:"CollateralType-currency"`
  AssessDate *string `xml:"assessDate,omitempty" key:"CollateralType-assessDate"`
  EncumbranceInd *string `xml:"encumbranceInd,omitempty" key:"CollateralType-encumbranceInd"`
  EndDate *string `xml:"endDate,omitempty" key:"CollateralType-endDate"`
  FactEndDate *string `xml:"factEndDate,omitempty" key:"CollateralType-factEndDate"`
  EndReason *string `xml:"endReason,omitempty" key:"CollateralType-endReason"`
}


// Блок 33. (юр Блок 24) Сведения о поручительстве
type XsWarrantyType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  ContractUid *string `xml:"contractUid,omitempty" key:"WarrantyType-contractUid"`
  Amount *string `xml:"amount,omitempty" key:"WarrantyType-amount"`
  Currency *string `xml:"currency,omitempty" key:"WarrantyType-currency"`
  StartDate *string `xml:"startDate,omitempty" key:"WarrantyType-startDate"`
  EndDate *string `xml:"endDate,omitempty" key:"WarrantyType-endDate"`
  FactEndDate *string `xml:"factEndDate,omitempty" key:"WarrantyType-factEndDate"`
  EndReason *string `xml:"endReason,omitempty" key:"WarrantyType-endReason"`
}


// Блок 34. (юр Блок 25) Сведения о независимой гарантии
type XsIndepGuaranteeType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  ContractUid *string `xml:"contractUid,omitempty" key:"IndepGuaranteeType-contractUid"`
  Amount *string `xml:"amount,omitempty" key:"IndepGuaranteeType-amount"`
  Currency *string `xml:"currency,omitempty" key:"IndepGuaranteeType-currency"`
  StartDate *string `xml:"startDate,omitempty" key:"IndepGuaranteeType-startDate"`
  EndDate *string `xml:"endDate,omitempty" key:"IndepGuaranteeType-endDate"`
  FactEndDate *string `xml:"factEndDate,omitempty" key:"IndepGuaranteeType-factEndDate"`
  EndReason *string `xml:"endReason,omitempty" key:"IndepGuaranteeType-endReason"`
}


// Блок 35. (юр Блок 26) Сведения о страховании предмета залога
type XsCollateralInsuranceType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  AssetId *string `xml:"assetId,omitempty" key:"CollateralInsuranceType-assetId"`
  Limit *string `xml:"limit,omitempty" key:"CollateralInsuranceType-limit"`
  Currency *string `xml:"currency,omitempty" key:"CollateralInsuranceType-currency"`
  FranchiseInd *string `xml:"franchiseInd,omitempty" key:"CollateralInsuranceType-franchiseInd"`
  StartDate *string `xml:"startDate,omitempty" key:"CollateralInsuranceType-startDate"`
  EndDate *string `xml:"endDate,omitempty" key:"CollateralInsuranceType-endDate"`
  FactEndDate *string `xml:"factEndDate,omitempty" key:"CollateralInsuranceType-factEndDate"`
  EndReason *string `xml:"endReason,omitempty" key:"CollateralInsuranceType-endReason"`
}


// Блок 36. (юр Блок 27) Сведения о погашении требований кредитора по обязательству за счет обеспечения
type XsSettlementType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  SecuringType *string `xml:"securingType,omitempty" key:"SettlementType-securingType"`
  Date *string `xml:"date,omitempty" key:"SettlementType-date"`
  Amount *string `xml:"amount,omitempty" key:"SettlementType-amount"`
}


// Блок 37. (юр Блок 28) Сведения о возмещении принципалом гаранту выплаченной суммы
type XsRepaymentType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  DueAmount *string `xml:"dueAmount,omitempty" key:"RepaymentType-dueAmount"`
  PaidAmount *string `xml:"paidAmount,omitempty" key:"RepaymentType-paidAmount"`
  ComplianceInd *string `xml:"complianceInd,omitempty" key:"RepaymentType-complianceInd"`
}


// Блок 38. (юр Блок 29) Сведения о прекращении обязательства
type XsTerminationType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Reason string `xml:"reason,omitempty" key:"TerminationType-reason"`
  Date string `xml:"date,omitempty" key:"TerminationType-date"`
}


// Блок 39. (юр Блок 30) Сведения о судебном споре или требовании по обязательству
type XsLawsuitType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  ActInd *string `xml:"actInd,omitempty" key:"LawsuitType-actInd"`
  ActDate *string `xml:"actDate,omitempty" key:"LawsuitType-actDate"`
  ActNum *string `xml:"actNum,omitempty" key:"LawsuitType-actNum"`
  ActResolution *string `xml:"actResolution,omitempty" key:"LawsuitType-actResolution"`
  EffectiveInd *string `xml:"effectiveInd,omitempty" key:"LawsuitType-effectiveInd"`
}


// Блок 40. Сведения квалифицированного бюро о среднемесячных платежах по договору займа (кредита)
type XsAvgPaymentType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Amount string `xml:"amount,omitempty" key:"AvgPaymentType-amount"`
  Date string `xml:"date,omitempty" key:"AvgPaymentType-date"`
  Currency string `xml:"currency,omitempty" key:"AvgPaymentType-currency"`
  DealUid string `xml:"dealUid,omitempty" key:"AvgPaymentType-dealUid"`
  BureauNum string `xml:"bureauNum,omitempty" key:"AvgPaymentType-bureauNum"`
}


// Блок 41. (юр Блок 31) Сведения об обязательстве, если в отношении источника открыто конкурсное производство
type XsSourceBankruptcyType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  StartOutstanding string `xml:"startOutstanding,omitempty" key:"SourceBankruptcyType-startOutstanding"`
  StartDate string `xml:"startDate,omitempty" key:"SourceBankruptcyType-startDate"`
  EndOutstanding *string `xml:"endOutstanding,omitempty" key:"SourceBankruptcyType-endOutstanding"`
  EndDate *string `xml:"endDate,omitempty" key:"SourceBankruptcyType-endDate"`
  LastPayOutstanding string `xml:"lastPayOutstanding,omitempty" key:"SourceBankruptcyType-lastPayOutstanding"`
  LastPayDate *string `xml:"lastPayDate,omitempty" key:"SourceBankruptcyType-lastPayDate"`
  OwnerChangeInd string `xml:"ownerChangeInd,omitempty" key:"SourceBankruptcyType-ownerChangeInd"`
  ObligationEndInd string `xml:"obligationEndInd,omitempty" key:"SourceBankruptcyType-obligationEndInd"`
  ObligationEndReason *string `xml:"obligationEndReason,omitempty" key:"SourceBankruptcyType-obligationEndReason"`
}


// Блок 42. (юр Блок 32) Сведения об обязательстве, если источник находится в процессе ликвидации
type XsSourceLiquidationType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  StartOutstanding string `xml:"startOutstanding,omitempty" key:"SourceLiquidationType-startOutstanding"`
  StartDate string `xml:"startDate,omitempty" key:"SourceLiquidationType-startDate"`
  EndOutstanding *string `xml:"endOutstanding,omitempty" key:"SourceLiquidationType-endOutstanding"`
  EndDate *string `xml:"endDate,omitempty" key:"SourceLiquidationType-endDate"`
  LastPayOutstanding string `xml:"lastPayOutstanding,omitempty" key:"SourceLiquidationType-lastPayOutstanding"`
  LastPayDate *string `xml:"lastPayDate,omitempty" key:"SourceLiquidationType-lastPayDate"`
  OwnerChangeInd string `xml:"ownerChangeInd,omitempty" key:"SourceLiquidationType-ownerChangeInd"`
  ObligationEndInd string `xml:"obligationEndInd,omitempty" key:"SourceLiquidationType-obligationEndInd"`
  ObligationEndReason *string `xml:"obligationEndReason,omitempty" key:"SourceLiquidationType-obligationEndReason"`
}


// Блок 43. (юр Блок 33) Сведения о взыскании долга по алиментам, платы за жилое помещение, коммунальные услуги или услуги связи
type XsCollectionType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  DebtType string `xml:"debtType,omitempty" key:"CollectionType-debtType"`
  Court string `xml:"court,omitempty" key:"CollectionType-court"`
  ActNum string `xml:"actNum,omitempty" key:"CollectionType-actNum"`
  ActDate string `xml:"actDate,omitempty" key:"CollectionType-actDate"`
  ActResolution string `xml:"actResolution,omitempty" key:"CollectionType-actResolution"`
  ExecNum *string `xml:"execNum,omitempty" key:"CollectionType-execNum"`
  CollectEndDate *string `xml:"collectEndDate,omitempty" key:"CollectionType-collectEndDate"`
  CollectOnceAmount *string `xml:"collectOnceAmount,omitempty" key:"CollectionType-collectOnceAmount"`
  CollectTotalAmount *string `xml:"collectTotalAmount,omitempty" key:"CollectionType-collectTotalAmount"`
  CalcDate *string `xml:"calcDate,omitempty" key:"CollectionType-calcDate"`
  CollectAmount *string `xml:"collectAmount,omitempty" key:"CollectionType-collectAmount"`
  Frequency string `xml:"frequency,omitempty" key:"CollectionType-frequency"`
}


// Блок 44. (юр Блок 34) Сведения о запросе информации пользователем
type XsCustomerRequestType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  InfoCode []string `xml:"infoCode,omitempty" key:"CustomerRequestType-infoCode"`
  ResultDate string `xml:"resultDate,omitempty" key:"CustomerRequestType-resultDate"`
  Date string `xml:"date,omitempty" key:"CustomerRequestType-date"`
  RequestReason []string `xml:"requestReason,omitempty" key:"CustomerRequestType-requestReason"`
  OtherReason *string `xml:"otherReason,omitempty" key:"CustomerRequestType-otherReason"`
  Amount *string `xml:"amount,omitempty" key:"CustomerRequestType-amount"`
  Currency *string `xml:"currency,omitempty" key:"CustomerRequestType-currency"`
}


// Блок 45. (юр Блок 35) Сведения о прекращении передачи информации по обязательству
type XsStopInfoType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Reason string `xml:"reason,omitempty" key:"StopInfoType-reason"`
  Date string `xml:"date,omitempty" key:"StopInfoType-date"`
}


// Блок 46. (юр Блок 36) Сведения об источнике – юридическом лице
type XsOrgSourceType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  RusRegInd string `xml:"rusRegInd,omitempty" key:"OrgSourceType-rusRegInd"`
  FullName string `xml:"fullName,omitempty" key:"OrgSourceType-fullName"`
  ShortName string `xml:"shortName,omitempty" key:"OrgSourceType-shortName"`
  OtherName *string `xml:"otherName,omitempty" key:"OrgSourceType-otherName"`
  Lei *string `xml:"lei,omitempty" key:"OrgSourceType-lei"`
  RegNum *string `xml:"regNum,omitempty" key:"OrgSourceType-regNum"`
  TaxCode *string `xml:"taxCode,omitempty" key:"OrgSourceType-taxCode"`
  TaxNum *string `xml:"taxNum,omitempty" key:"OrgSourceType-taxNum"`
  SourceCode string `xml:"sourceCode,omitempty" key:"OrgSourceType-sourceCode"`
  CreationDate string `xml:"creationDate,omitempty" key:"OrgSourceType-creationDate"`
  BankruptcyStartDate *string `xml:"bankruptcyStartDate,omitempty" key:"OrgSourceType-bankruptcyStartDate"`
  BankruptcyEndDate *string `xml:"bankruptcyEndDate,omitempty" key:"OrgSourceType-bankruptcyEndDate"`
  LiquidationStartDate *string `xml:"liquidationStartDate,omitempty" key:"OrgSourceType-liquidationStartDate"`
  LiquidationEndDate *string `xml:"liquidationEndDate,omitempty" key:"OrgSourceType-liquidationEndDate"`
}


// Блок 47. (юр Блок 37) Сведения об источнике – физическом лице
type XsPersonSourceType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  FirstName string `xml:"firstName,omitempty" key:"PersonSourceType-firstName"`
  LastName string `xml:"lastName,omitempty" key:"PersonSourceType-lastName"`
  MidName *string `xml:"midName,omitempty" key:"PersonSourceType-midName"`
  BirthDate string `xml:"birthDate,omitempty" key:"PersonSourceType-birthDate"`
  BirthPlace string `xml:"birthPlace,omitempty" key:"PersonSourceType-birthPlace"`
  IdCode string `xml:"idCode,omitempty" key:"PersonSourceType-idCode"`
  OtherId *string `xml:"otherId,omitempty" key:"PersonSourceType-otherId"`
  IdSeries *string `xml:"idSeries,omitempty" key:"PersonSourceType-idSeries"`
  IdNum string `xml:"idNum,omitempty" key:"PersonSourceType-idNum"`
  IssueDate string `xml:"issueDate,omitempty" key:"PersonSourceType-issueDate"`
  Issuer string `xml:"issuer,omitempty" key:"PersonSourceType-issuer"`
  DeptCode *string `xml:"deptCode,omitempty" key:"PersonSourceType-deptCode"`
  RegNum *string `xml:"regNum,omitempty" key:"PersonSourceType-regNum"`
}


// Блок 48. (юр Блок 38) Сведения об источнике – арбитражном управляющем
type XsCommissionerSourceType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  FirstName string `xml:"firstName,omitempty" key:"CommissionerSourceType-firstName"`
  LastName string `xml:"lastName,omitempty" key:"CommissionerSourceType-lastName"`
  MidName *string `xml:"midName,omitempty" key:"CommissionerSourceType-midName"`
  BirthDate string `xml:"birthDate,omitempty" key:"CommissionerSourceType-birthDate"`
  BirthPlace string `xml:"birthPlace,omitempty" key:"CommissionerSourceType-birthPlace"`
  OrgName *string `xml:"orgName,omitempty" key:"CommissionerSourceType-orgName"`
  OrgAddress string `xml:"orgAddress,omitempty" key:"CommissionerSourceType-orgAddress"`
  StartDate string `xml:"startDate,omitempty" key:"CommissionerSourceType-startDate"`
  EndDate *string `xml:"endDate,omitempty" key:"CommissionerSourceType-endDate"`
  TaxCode *string `xml:"taxCode,omitempty" key:"CommissionerSourceType-taxCode"`
  TaxNum *string `xml:"taxNum,omitempty" key:"CommissionerSourceType-taxNum"`
  SocialId string `xml:"socialId,omitempty" key:"CommissionerSourceType-socialId"`
}


// Блок 49. (юр Блок 39) Сведения о пользователе – юридическом лице
type XsOrgCustomerType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  RusRegInd string `xml:"rusRegInd,omitempty" key:"OrgCustomerType-rusRegInd"`
  FullName string `xml:"fullName,omitempty" key:"OrgCustomerType-fullName"`
  ShortName string `xml:"shortName,omitempty" key:"OrgCustomerType-shortName"`
  OtherName *string `xml:"otherName,omitempty" key:"OrgCustomerType-otherName"`
  Lei *string `xml:"lei,omitempty" key:"OrgCustomerType-lei"`
  RegNum *string `xml:"regNum,omitempty" key:"OrgCustomerType-regNum"`
  TaxCode *string `xml:"taxCode,omitempty" key:"OrgCustomerType-taxCode"`
  TaxNum *string `xml:"taxNum,omitempty" key:"OrgCustomerType-taxNum"`
  CustomerCode string `xml:"customerCode,omitempty" key:"OrgCustomerType-customerCode"`
  MonitoringInd string `xml:"monitoringInd,omitempty" key:"OrgCustomerType-monitoringInd"`
  MonitoringDate *string `xml:"monitoringDate,omitempty" key:"OrgCustomerType-monitoringDate"`
  InfoCode []string `xml:"infoCode,omitempty" key:"OrgCustomerType-infoCode"`
  ResultDate string `xml:"resultDate,omitempty" key:"OrgCustomerType-resultDate"`
  Date string `xml:"date,omitempty" key:"OrgCustomerType-date"`
  RequestReason []string `xml:"requestReason,omitempty" key:"OrgCustomerType-requestReason"`
  OtherReason *string `xml:"otherReason,omitempty" key:"OrgCustomerType-otherReason"`
  Amount *string `xml:"amount,omitempty" key:"OrgCustomerType-amount"`
  Currency *string `xml:"currency,omitempty" key:"OrgCustomerType-currency"`
}


// Блок 50. (юр Блок 40) Сведения о пользователе – индивидуальном предпринимателе
type XsPersonCustomerType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  FirstName string `xml:"firstName,omitempty" key:"PersonCustomerType-firstName"`
  LastName string `xml:"lastName,omitempty" key:"PersonCustomerType-lastName"`
  MidName *string `xml:"midName,omitempty" key:"PersonCustomerType-midName"`
  BirthDate string `xml:"birthDate,omitempty" key:"PersonCustomerType-birthDate"`
  BirthPlace string `xml:"birthPlace,omitempty" key:"PersonCustomerType-birthPlace"`
  IdCode string `xml:"idCode,omitempty" key:"PersonCustomerType-idCode"`
  OtherId *string `xml:"otherId,omitempty" key:"PersonCustomerType-otherId"`
  IdSeries *string `xml:"idSeries,omitempty" key:"PersonCustomerType-idSeries"`
  IdNum string `xml:"idNum,omitempty" key:"PersonCustomerType-idNum"`
  IssueDate string `xml:"issueDate,omitempty" key:"PersonCustomerType-issueDate"`
  Issuer string `xml:"issuer,omitempty" key:"PersonCustomerType-issuer"`
  DeptCode *string `xml:"deptCode,omitempty" key:"PersonCustomerType-deptCode"`
  MonitoringInd string `xml:"monitoringInd,omitempty" key:"PersonCustomerType-monitoringInd"`
  MonitoringDate *string `xml:"monitoringDate,omitempty" key:"PersonCustomerType-monitoringDate"`
  InfoCode []string `xml:"infoCode,omitempty" key:"PersonCustomerType-infoCode"`
  ResultDate string `xml:"resultDate,omitempty" key:"PersonCustomerType-resultDate"`
  Date string `xml:"date,omitempty" key:"PersonCustomerType-date"`
  RequestReason []string `xml:"requestReason,omitempty" key:"PersonCustomerType-requestReason"`
  OtherReason *string `xml:"otherReason,omitempty" key:"PersonCustomerType-otherReason"`
  Amount *string `xml:"amount,omitempty" key:"PersonCustomerType-amount"`
  Currency *string `xml:"currency,omitempty" key:"PersonCustomerType-currency"`
  TaxNum *string `xml:"taxNum,omitempty" key:"PersonCustomerType-taxNum"`
  TaxCode *string `xml:"taxCode,omitempty" key:"PersonCustomerType-taxCode"`
  RegNum *string `xml:"regNum,omitempty" key:"PersonCustomerType-regNum"`
  SocialId string `xml:"socialId,omitempty" key:"PersonCustomerType-socialId"`
}


// Блок 51. (юр Блок 41) Сведения о приобретателе прав – юридическом лице
type XsOrgAcquirerInfoType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  RusRegInd string `xml:"rusRegInd,omitempty" key:"OrgAcquirerInfoType-rusRegInd"`
  FullName string `xml:"fullName,omitempty" key:"OrgAcquirerInfoType-fullName"`
  ShortName string `xml:"shortName,omitempty" key:"OrgAcquirerInfoType-shortName"`
  OtherName *string `xml:"otherName,omitempty" key:"OrgAcquirerInfoType-otherName"`
  Lei *string `xml:"lei,omitempty" key:"OrgAcquirerInfoType-lei"`
  RegNum *string `xml:"regNum,omitempty" key:"OrgAcquirerInfoType-regNum"`
  TaxCode *string `xml:"taxCode,omitempty" key:"OrgAcquirerInfoType-taxCode"`
  TaxNum *string `xml:"taxNum,omitempty" key:"OrgAcquirerInfoType-taxNum"`
  AcquirerCode string `xml:"acquirerCode,omitempty" key:"OrgAcquirerInfoType-acquirerCode"`
  AcquireDate string `xml:"acquireDate,omitempty" key:"OrgAcquirerInfoType-acquireDate"`
}


// Блок 52. (юр Блок 42) Сведения о приобретателе прав – физическом лице
type XsPersonAcquirerInfoType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  FirstName string `xml:"firstName,omitempty" key:"PersonAcquirerInfoType-firstName"`
  LastName string `xml:"lastName,omitempty" key:"PersonAcquirerInfoType-lastName"`
  MidName *string `xml:"midName,omitempty" key:"PersonAcquirerInfoType-midName"`
  BirthDate string `xml:"birthDate,omitempty" key:"PersonAcquirerInfoType-birthDate"`
  BirthPlace string `xml:"birthPlace,omitempty" key:"PersonAcquirerInfoType-birthPlace"`
  IdCode string `xml:"idCode,omitempty" key:"PersonAcquirerInfoType-idCode"`
  OtherId *string `xml:"otherId,omitempty" key:"PersonAcquirerInfoType-otherId"`
  IdSeries *string `xml:"idSeries,omitempty" key:"PersonAcquirerInfoType-idSeries"`
  IdNum string `xml:"idNum,omitempty" key:"PersonAcquirerInfoType-idNum"`
  IssueDate string `xml:"issueDate,omitempty" key:"PersonAcquirerInfoType-issueDate"`
  Issuer string `xml:"issuer,omitempty" key:"PersonAcquirerInfoType-issuer"`
  DeptCode *string `xml:"deptCode,omitempty" key:"PersonAcquirerInfoType-deptCode"`
  AcquireDate string `xml:"acquireDate,omitempty" key:"PersonAcquirerInfoType-acquireDate"`
  TaxCode *string `xml:"taxCode,omitempty" key:"PersonAcquirerInfoType-taxCode"`
  TaxNum *string `xml:"taxNum,omitempty" key:"PersonAcquirerInfoType-taxNum"`
  SocialId string `xml:"socialId,omitempty" key:"PersonAcquirerInfoType-socialId"`
}


// Блок 53. (юр Блок 43) Сведения об обслуживающей организации
type XsServiceOrgType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  RusRegInd string `xml:"rusRegInd,omitempty" key:"ServiceOrgType-rusRegInd"`
  FullName string `xml:"fullName,omitempty" key:"ServiceOrgType-fullName"`
  ShortName string `xml:"shortName,omitempty" key:"ServiceOrgType-shortName"`
  OtherName *string `xml:"otherName,omitempty" key:"ServiceOrgType-otherName"`
  RegNum *string `xml:"regNum,omitempty" key:"ServiceOrgType-regNum"`
  TaxCode *string `xml:"taxCode,omitempty" key:"ServiceOrgType-taxCode"`
  TaxNum *string `xml:"taxNum,omitempty" key:"ServiceOrgType-taxNum"`
  StartDate *string `xml:"startDate,omitempty" key:"ServiceOrgType-startDate"`
  EndDate *string `xml:"endDate,omitempty" key:"ServiceOrgType-endDate"`
  IssuerName *string `xml:"issuerName,omitempty" key:"ServiceOrgType-issuerName"`
  IssuerRegNum string `xml:"issuerRegNum,omitempty" key:"ServiceOrgType-issuerRegNum"`
}


// Блок 54. (юр Блок 44) Сведения об учете обязательства
type XsAccountingType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  AccountingInd string `xml:"accountingInd,omitempty" key:"AccountingType-accountingInd"`
}


// Блок 55. (юр Блок 45) Сведения об обращении субъекта к источнику с предложением совершить сделку
type XsApplicationType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"ApplicationType-subjectRole"`
  Amount *string `xml:"amount,omitempty" key:"ApplicationType-amount"`
  Currency *string `xml:"currency,omitempty" key:"ApplicationType-currency"`
  Uid string `xml:"uid,omitempty" key:"ApplicationType-uid"`
  Date string `xml:"date,omitempty" key:"ApplicationType-date"`
  SourceKind string `xml:"sourceKind,omitempty" key:"ApplicationType-sourceKind"`
  SubmitKind string `xml:"submitKind,omitempty" key:"ApplicationType-submitKind"`
  ExpireDate string `xml:"expireDate,omitempty" key:"ApplicationType-expireDate"`
}


// Блок 56. (юр Блок 46) Сведения об участии в обязательстве, по которому формируется кредитная история
type XsParticipationType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  SubjectRole string `xml:"subjectRole,omitempty" key:"ParticipationType-subjectRole"`
  LoanKind *string `xml:"loanKind,omitempty" key:"ParticipationType-loanKind"`
  DealUid string `xml:"dealUid,omitempty" key:"ParticipationType-dealUid"`
  TransferDate string `xml:"transferDate,omitempty" key:"ParticipationType-transferDate"`
  DefaultInd string `xml:"defaultInd,omitempty" key:"ParticipationType-defaultInd"`
  StopInd string `xml:"stopInd,omitempty" key:"ParticipationType-stopInd"`
}


// Блок 57. (юр Блок 47) Сведения об отказе источника от предложения совершить сделку
type XsAppRejectType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Date string `xml:"date,omitempty" key:"AppRejectType-date"`
  Reason []string `xml:"reason,omitempty" key:"AppRejectType-reason"`
}


// Титульная часть КИ Субъекта, используется только для поиска Субъекта
type XsPersonSubjectType struct {
  Name XsPersonNameType `xml:"name,omitempty" key:"PersonSubjectType-name"`
  PrevName []XsPersonPrevNameType `xml:"prevName,omitempty" key:"PersonSubjectType-prevName"`
  BirthInfo XsBirthInfoType `xml:"birthInfo,omitempty" key:"PersonSubjectType-birthInfo"`
  Id []XsPersonIdType `xml:"id,omitempty" key:"PersonSubjectType-id"`
  PrevId []XsPersonPrevIdType `xml:"prevId,omitempty" key:"PersonSubjectType-prevId"`
  Tax []XsTaxType `xml:"tax,omitempty" key:"PersonSubjectType-tax"`
  SocialId *XsSocialIdType `xml:"socialId,omitempty" key:"PersonSubjectType-socialId"`
}


// Базовый общий блок для источника - физического лица
type XsPersonSourceCommonType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  FirstName string `xml:"firstName,omitempty" key:"PersonSourceCommonType-firstName"`
  LastName string `xml:"lastName,omitempty" key:"PersonSourceCommonType-lastName"`
  MidName *string `xml:"midName,omitempty" key:"PersonSourceCommonType-midName"`
  BirthDate string `xml:"birthDate,omitempty" key:"PersonSourceCommonType-birthDate"`
  BirthPlace string `xml:"birthPlace,omitempty" key:"PersonSourceCommonType-birthPlace"`
  IdCode string `xml:"idCode,omitempty" key:"PersonSourceCommonType-idCode"`
  OtherId *string `xml:"otherId,omitempty" key:"PersonSourceCommonType-otherId"`
  IdSeries *string `xml:"idSeries,omitempty" key:"PersonSourceCommonType-idSeries"`
  IdNum string `xml:"idNum,omitempty" key:"PersonSourceCommonType-idNum"`
  IssueDate string `xml:"issueDate,omitempty" key:"PersonSourceCommonType-issueDate"`
  Issuer string `xml:"issuer,omitempty" key:"PersonSourceCommonType-issuer"`
  DeptCode *string `xml:"deptCode,omitempty" key:"PersonSourceCommonType-deptCode"`
}


// Базовый общий блок для источника - юридического лица
type XsOrgSourceCommonType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  RusRegInd string `xml:"rusRegInd,omitempty" key:"OrgSourceCommonType-rusRegInd"`
  FullName string `xml:"fullName,omitempty" key:"OrgSourceCommonType-fullName"`
  ShortName string `xml:"shortName,omitempty" key:"OrgSourceCommonType-shortName"`
  OtherName *string `xml:"otherName,omitempty" key:"OrgSourceCommonType-otherName"`
  Lei *string `xml:"lei,omitempty" key:"OrgSourceCommonType-lei"`
  RegNum *string `xml:"regNum,omitempty" key:"OrgSourceCommonType-regNum"`
  TaxCode *string `xml:"taxCode,omitempty" key:"OrgSourceCommonType-taxCode"`
  TaxNum *string `xml:"taxNum,omitempty" key:"OrgSourceCommonType-taxNum"`
}


// Кредитная история юридического лица
type XsOrgCreditHistoryType struct {
  Name XsOrgNameType `xml:"name,omitempty" key:"OrgCreditHistoryType-name"`
  Address XsOrgAddressType `xml:"address,omitempty" key:"OrgCreditHistoryType-address"`
  RegNum XsOrgRegNumType `xml:"regNum,omitempty" key:"OrgCreditHistoryType-regNum"`
  Tax *XsOrgTaxType `xml:"tax,omitempty" key:"OrgCreditHistoryType-tax"`
  OrgChange []XsOrgChangeType `xml:"orgChange,omitempty" key:"OrgCreditHistoryType-orgChange"`
  Bankruptcy []XsBankruptcyType `xml:"bankruptcy,omitempty" key:"OrgCreditHistoryType-bankruptcy"`
  Fulfillment []XsOrgFulfillmentType `xml:"fulfillment,omitempty" key:"OrgCreditHistoryType-fulfillment"`
  PrevOrg []XsPrevOrgType `xml:"prevOrg,omitempty" key:"OrgCreditHistoryType-prevOrg"`
  Score []XsScoreType `xml:"score,omitempty" key:"OrgCreditHistoryType-score"`
  DealUid XsDealUidType `xml:"dealUid,omitempty" key:"OrgCreditHistoryType-dealUid"`
  Main XsOrgDealMainType `xml:"main,omitempty" key:"OrgCreditHistoryType-main"`
  Amount *XsDealAmountType `xml:"amount,omitempty" key:"OrgCreditHistoryType-amount"`
  JointDebtors *XsJointDebtorsType `xml:"jointDebtors,omitempty" key:"OrgCreditHistoryType-jointDebtors"`
  PaymentTerms *XsPaymentTermsType `xml:"paymentTerms,omitempty" key:"OrgCreditHistoryType-paymentTerms"`
  Change []XsDealChangeType `xml:"change,omitempty" key:"OrgCreditHistoryType-change"`
  FundDate *XsFundDateType `xml:"fundDate,omitempty" key:"OrgCreditHistoryType-fundDate"`
  Arrear []XsArrearType `xml:"arrear,omitempty" key:"OrgCreditHistoryType-arrear"`
  DueArrear []XsDueArrearType `xml:"dueArrear,omitempty" key:"OrgCreditHistoryType-dueArrear"`
  PastdueArrear []XsPastdueArrearType `xml:"pastdueArrear,omitempty" key:"OrgCreditHistoryType-pastdueArrear"`
  Payments []XsPaymentsType `xml:"payments,omitempty" key:"OrgCreditHistoryType-payments"`
  MonthlyPayment *XsMonthlyPaymentType `xml:"monthlyPayment,omitempty" key:"OrgCreditHistoryType-monthlyPayment"`
  SourceNonMonetaryObligation *XsSourceNonMonetaryObligationType `xml:"sourceNonMonetaryObligation,omitempty" key:"OrgCreditHistoryType-sourceNonMonetaryObligation"`
  SubjectNonMonetaryObligation *XsSubjectNonMonetaryObligationType `xml:"subjectNonMonetaryObligation,omitempty" key:"OrgCreditHistoryType-subjectNonMonetaryObligation"`
  Collateral []XsCollateralType `xml:"collateral,omitempty" key:"OrgCreditHistoryType-collateral"`
  Warranty []XsWarrantyType `xml:"warranty,omitempty" key:"OrgCreditHistoryType-warranty"`
  IndepGuarantee []XsIndepGuaranteeType `xml:"indepGuarantee,omitempty" key:"OrgCreditHistoryType-indepGuarantee"`
  CollateralInsurance []XsCollateralInsuranceType `xml:"collateralInsurance,omitempty" key:"OrgCreditHistoryType-collateralInsurance"`
  Settlement []XsSettlementType `xml:"settlement,omitempty" key:"OrgCreditHistoryType-settlement"`
  Repayment *XsRepaymentType `xml:"repayment,omitempty" key:"OrgCreditHistoryType-repayment"`
  Termination *XsTerminationType `xml:"termination,omitempty" key:"OrgCreditHistoryType-termination"`
  Lawsuit []XsLawsuitType `xml:"lawsuit,omitempty" key:"OrgCreditHistoryType-lawsuit"`
  SourceBankruptcy *XsSourceBankruptcyType `xml:"sourceBankruptcy,omitempty" key:"OrgCreditHistoryType-sourceBankruptcy"`
  SourceLiquidation *XsSourceLiquidationType `xml:"sourceLiquidation,omitempty" key:"OrgCreditHistoryType-sourceLiquidation"`
  Collection *XsCollectionType `xml:"collection,omitempty" key:"OrgCreditHistoryType-collection"`
  InfoDate string `xml:"infoDate,omitempty" key:"OrgCreditHistoryType-infoDate"`
  OrgSource *XsOrgSourceType `xml:"orgSource,omitempty" key:"OrgCreditHistoryType-orgSource"`
  PersonSource *XsPersonSourceType `xml:"personSource,omitempty" key:"OrgCreditHistoryType-personSource"`
  CommissionerSource *XsCommissionerSourceType `xml:"commissionerSource,omitempty" key:"OrgCreditHistoryType-commissionerSource"`
  OrgCustomer []XsOrgCustomerType `xml:"orgCustomer,omitempty" key:"OrgCreditHistoryType-orgCustomer"`
  PersonCustomer []XsPersonCustomerType `xml:"personCustomer,omitempty" key:"OrgCreditHistoryType-personCustomer"`
  OrgAcquirerInfo *XsOrgAcquirerInfoType `xml:"orgAcquirerInfo,omitempty" key:"OrgCreditHistoryType-orgAcquirerInfo"`
  PersonAcquirerInfo *XsPersonAcquirerInfoType `xml:"personAcquirerInfo,omitempty" key:"OrgCreditHistoryType-personAcquirerInfo"`
  ServiceOrg *XsServiceOrgType `xml:"serviceOrg,omitempty" key:"OrgCreditHistoryType-serviceOrg"`
  Accounting *XsAccountingType `xml:"accounting,omitempty" key:"OrgCreditHistoryType-accounting"`
  Application []XsApplicationType `xml:"application,omitempty" key:"OrgCreditHistoryType-application"`
  Participation []XsParticipationType `xml:"participation,omitempty" key:"OrgCreditHistoryType-participation"`
  AppReject []XsAppRejectType `xml:"appReject,omitempty" key:"OrgCreditHistoryType-appReject"`
}


// юр Блок 1. Наименование юридического лица
type XsOrgNameType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  FullName string `xml:"fullName,omitempty" key:"OrgNameType-fullName"`
  ShortName string `xml:"shortName,omitempty" key:"OrgNameType-shortName"`
  OtherName *string `xml:"otherName,omitempty" key:"OrgNameType-otherName"`
}


// юр Блок 2. Адрес юридического лица в пределах его места нахождения и контактная информация
type XsOrgAddressType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  CountryCode string `xml:"countryCode,omitempty" key:"OrgAddressType-countryCode"`
  OtherCountry *string `xml:"otherCountry,omitempty" key:"OrgAddressType-otherCountry"`
  AddrNum *string `xml:"addrNum,omitempty" key:"OrgAddressType-addrNum"`
  LocationCode *string `xml:"locationCode,omitempty" key:"OrgAddressType-locationCode"`
  OtherLocation *string `xml:"otherLocation,omitempty" key:"OrgAddressType-otherLocation"`
  Street *string `xml:"street,omitempty" key:"OrgAddressType-street"`
  House *string `xml:"house,omitempty" key:"OrgAddressType-house"`
  Estate *string `xml:"estate,omitempty" key:"OrgAddressType-estate"`
  Block *string `xml:"block,omitempty" key:"OrgAddressType-block"`
  Building *string `xml:"building,omitempty" key:"OrgAddressType-building"`
  Office *string `xml:"office,omitempty" key:"OrgAddressType-office"`
  Phone *string `xml:"phone,omitempty" key:"OrgAddressType-phone"`
  PhoneComment *string `xml:"phoneComment,omitempty" key:"OrgAddressType-phoneComment"`
  Email *string `xml:"email,omitempty" key:"OrgAddressType-email"`
}


// юр Блок 3. Регистрационный номер
type XsOrgRegNumType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Num *string `xml:"num,omitempty" key:"OrgRegNumType-num"`
  Lei *string `xml:"lei,omitempty" key:"OrgRegNumType-lei"`
}


// юр Блок 4. Номер налогоплательщика
type XsOrgTaxType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  TaxCode *string `xml:"taxCode,omitempty" key:"OrgTaxType-taxCode"`
  TaxNum *string `xml:"taxNum,omitempty" key:"OrgTaxType-taxNum"`
}


// юр Блок 5. Сведения о смене наименования либо правопреемстве при реорганизации
type XsOrgChangeType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  RenameInd string `xml:"renameInd,omitempty" key:"OrgChangeType-renameInd"`
  PrevShortName *string `xml:"prevShortName,omitempty" key:"OrgChangeType-prevShortName"`
  PrevFullName *string `xml:"prevFullName,omitempty" key:"OrgChangeType-prevFullName"`
  ReorgInd string `xml:"reorgInd,omitempty" key:"OrgChangeType-reorgInd"`
  PrevRegNum *string `xml:"prevRegNum,omitempty" key:"OrgChangeType-prevRegNum"`
  Date *string `xml:"date,omitempty" key:"OrgChangeType-date"`
}


// юр Блок 7. Сведения о завершении расчетов с кредиторами и освобождении субъекта от исполнения обязательств в связи с банкротством
type XsOrgFulfillmentType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Date *string `xml:"date,omitempty" key:"OrgFulfillmentType-date"`
}


// юр Блок 8. Сведения об основных частях кредитных историй юридического лица, от которого субъекту перешли права и обязанности
type XsPrevOrgType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  RusRegInd *string `xml:"rusRegInd,omitempty" key:"PrevOrgType-rusRegInd"`
  Name *string `xml:"name,omitempty" key:"PrevOrgType-name"`
  RegNum *string `xml:"regNum,omitempty" key:"PrevOrgType-regNum"`
}


// юр Блок 11. Общие сведения о сделке
type XsOrgDealMainType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
  Date string `xml:"date,omitempty" key:"OrgDealMainType-date"`
  DealType string `xml:"dealType,omitempty" key:"OrgDealMainType-dealType"`
  LoanKind *string `xml:"loanKind,omitempty" key:"OrgDealMainType-loanKind"`
  LoanPurpose []string `xml:"loanPurpose,omitempty" key:"OrgDealMainType-loanPurpose"`
  CardInd string `xml:"cardInd,omitempty" key:"OrgDealMainType-cardInd"`
  NovationInd string `xml:"novationInd,omitempty" key:"OrgDealMainType-novationInd"`
  SourceObligationInd *string `xml:"sourceObligationInd,omitempty" key:"OrgDealMainType-sourceObligationInd"`
  SubjectObligationInd *string `xml:"subjectObligationInd,omitempty" key:"OrgDealMainType-subjectObligationInd"`
  EndDate string `xml:"endDate,omitempty" key:"OrgDealMainType-endDate"`
}


// Титульная часть КИ Субъекта, используется только для поиска Субъекта
type XsOrgSubjectType struct {
  Name XsOrgNameType `xml:"name,omitempty" key:"OrgSubjectType-name"`
  Address XsOrgAddressType `xml:"address,omitempty" key:"OrgSubjectType-address"`
  RegNum XsOrgRegNumType `xml:"regNum,omitempty" key:"OrgSubjectType-regNum"`
  Tax *XsOrgTaxType `xml:"tax,omitempty" key:"OrgSubjectType-tax"`
  OrgChange []XsOrgChangeType `xml:"orgChange,omitempty" key:"OrgSubjectType-orgChange"`
}


// Базовый тип для всех Блоков (абстрактный)
type XsBlockType struct {
  InfoDate *string `xml:"infoDate,attr,omitempty"`
}


