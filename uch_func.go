// Auto-generated code - no manual modifications
package ucbcore

// Пакет данных
type XsPackageTypeFunc func(this *XsPackageType, ctx *ParseContext) error

var XsPackageTypeFuncInit XsPackageTypeFunc = (*XsPackageType).Init

var XsPackageTypeFuncLoad XsPackageTypeFunc = (*XsPackageType).Load

var XsPackageTypeFuncValidate XsPackageTypeFunc = (*XsPackageType).Validate

func (this *XsPackageType) Init(ctx *ParseContext) error {
  if err := XsSourceTypeFuncInit(&this.Source, ctx); err != nil { return err }
  for _, v := range (*this).Record {
    if err := XsRecordTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsPackageType) Load(ctx *ParseContext) error {
  if err := XsSourceTypeFuncLoad(&this.Source, ctx); err != nil { return err }
  for _, v := range (*this).Record {
    if err := XsRecordTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsPackageType) Validate(ctx *ParseContext) error {
  if err := XsSourceTypeFuncValidate(&this.Source, ctx); err != nil { return err }
  for _, v := range (*this).Record {
    if err := XsRecordTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// Запись в пакете
type XsRecordTypeFunc func(this *XsRecordType, ctx *ParseContext) error

var XsRecordTypeFuncInit XsRecordTypeFunc = (*XsRecordType).Init

var XsRecordTypeFuncLoad XsRecordTypeFunc = (*XsRecordType).Load

var XsRecordTypeFuncValidate XsRecordTypeFunc = (*XsRecordType).Validate

func (this *XsRecordType) Init(ctx *ParseContext) error {
  if this.PersonEvent != nil {
    if err := XsPersonEventTypeFuncInit(this.PersonEvent, ctx); err != nil { return err }
  }
  if this.OrgEvent != nil {
    if err := XsOrgEventTypeFuncInit(this.OrgEvent, ctx); err != nil { return err }
  }
  if this.InternalEvent != nil {
    if err := XsInternalEventTypeFuncInit(this.InternalEvent, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsRecordType) Load(ctx *ParseContext) error {
  if this.PersonEvent != nil {
    if err := XsPersonEventTypeFuncLoad(this.PersonEvent, ctx); err != nil { return err }
  }
  if this.OrgEvent != nil {
    if err := XsOrgEventTypeFuncLoad(this.OrgEvent, ctx); err != nil { return err }
  }
  if this.InternalEvent != nil {
    if err := XsInternalEventTypeFuncLoad(this.InternalEvent, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsRecordType) Validate(ctx *ParseContext) error {
  if this.PersonEvent != nil {
    if err := XsPersonEventTypeFuncValidate(this.PersonEvent, ctx); err != nil { return err }
  }
  if this.OrgEvent != nil {
    if err := XsOrgEventTypeFuncValidate(this.OrgEvent, ctx); err != nil { return err }
  }
  if this.InternalEvent != nil {
    if err := XsInternalEventTypeFuncValidate(this.InternalEvent, ctx); err != nil { return err }
  }
  return nil
}

// Событие Кредитной Истории Физического Лица
type XsPersonEventTypeFunc func(this *XsPersonEventType, ctx *ParseContext) error

var XsPersonEventTypeFuncInit XsPersonEventTypeFunc = (*XsPersonEventType).Init

var XsPersonEventTypeFuncLoad XsPersonEventTypeFunc = (*XsPersonEventType).Load

var XsPersonEventTypeFuncValidate XsPersonEventTypeFunc = (*XsPersonEventType).Validate

func (this *XsPersonEventType) Init(ctx *ParseContext) error {
  if err := XsPersonSubjectTypeFuncInit(&this.Subject, ctx); err != nil { return err }
  if this.SubjectApplied != nil {
    if err := XsSubjectAppliedTypeFuncInit(this.SubjectApplied, ctx); err != nil { return err }
  }
  if this.ApplicationChanged != nil {
    if err := XsApplicationChangedTypeFuncInit(this.ApplicationChanged, ctx); err != nil { return err }
  }
  if this.ApplicationRejected != nil {
    if err := XsApplicationRejectedTypeFuncInit(this.ApplicationRejected, ctx); err != nil { return err }
  }
  if this.MonetaryDealMade != nil {
    if err := XsMonetaryDealMadeTypeFuncInit(this.MonetaryDealMade, ctx); err != nil { return err }
  }
  if this.SourceNonMonetaryDealMade != nil {
    if err := XsSourceNonMonetaryDealMadeTypeFuncInit(this.SourceNonMonetaryDealMade, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealMade != nil {
    if err := XsNonMonetaryDealMadeTypeFuncInit(this.NonMonetaryDealMade, ctx); err != nil { return err }
  }
  if this.LeasingDealMade != nil {
    if err := XsLeasingDealMadeTypeFuncInit(this.LeasingDealMade, ctx); err != nil { return err }
  }
  if this.DebtClaimSubmitted != nil {
    if err := XsDebtClaimSubmittedTypeFuncInit(this.DebtClaimSubmitted, ctx); err != nil { return err }
  }
  if this.DebtClaimChanged != nil {
    if err := XsDebtClaimChangedTypeFuncInit(this.DebtClaimChanged, ctx); err != nil { return err }
  }
  if this.SubjectMainChanged != nil {
    if err := XsSubjectMainChangedTypeFuncInit(this.SubjectMainChanged, ctx); err != nil { return err }
  }
  if this.SubjectSpecialChanged != nil {
    if err := XsSubjectSpecialChangedTypeFuncInit(this.SubjectSpecialChanged, ctx); err != nil { return err }
  }
  if this.SubjectCapacityChanged != nil {
    if err := XsSubjectCapacityChangedTypeFuncInit(this.SubjectCapacityChanged, ctx); err != nil { return err }
  }
  if this.SubjectBankruptcyChanged != nil {
    if err := XsSubjectBankruptcyChangedTypeFuncInit(this.SubjectBankruptcyChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealChanged != nil {
    if err := XsMonetaryDealChangedTypeFuncInit(this.MonetaryDealChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealChanged != nil {
    if err := XsNonMonetaryDealChangedTypeFuncInit(this.NonMonetaryDealChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealPrincipalChanged != nil {
    if err := XsMonetaryDealPrincipalChangedTypeFuncInit(this.MonetaryDealPrincipalChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealPrincipalChanged != nil {
    if err := XsNonMonetaryDealPrincipalChangedTypeFuncInit(this.NonMonetaryDealPrincipalChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealPerformanceChanged != nil {
    if err := XsMonetaryDealPerformanceChangedTypeFuncInit(this.MonetaryDealPerformanceChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealPerformanceChanged != nil {
    if err := XsNonMonetaryDealPerformanceChangedTypeFuncInit(this.NonMonetaryDealPerformanceChanged, ctx); err != nil { return err }
  }
  if this.DealSecuringChanged != nil {
    if err := XsDealSecuringChangedTypeFuncInit(this.DealSecuringChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealEnded != nil {
    if err := XsMonetaryDealEndedTypeFuncInit(this.MonetaryDealEnded, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealEnded != nil {
    if err := XsNonMonetaryDealEndedTypeFuncInit(this.NonMonetaryDealEnded, ctx); err != nil { return err }
  }
  if this.DealClaimChanged != nil {
    if err := XsDealClaimChangedTypeFuncInit(this.DealClaimChanged, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyStarted != nil {
    if err := XsSourceBankruptcyStartedTypeFuncInit(this.SourceBankruptcyStarted, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyChanged != nil {
    if err := XsSourceBankruptcyChangedTypeFuncInit(this.SourceBankruptcyChanged, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyEnded != nil {
    if err := XsSourceBankruptcyEndedTypeFuncInit(this.SourceBankruptcyEnded, ctx); err != nil { return err }
  }
  if this.SourceLiquidationStarted != nil {
    if err := XsSourceLiquidationStartedTypeFuncInit(this.SourceLiquidationStarted, ctx); err != nil { return err }
  }
  if this.SourceLiquidationChanged != nil {
    if err := XsSourceLiquidationChangedTypeFuncInit(this.SourceLiquidationChanged, ctx); err != nil { return err }
  }
  if this.SourceLiquidationEnded != nil {
    if err := XsSourceLiquidationEndedTypeFuncInit(this.SourceLiquidationEnded, ctx); err != nil { return err }
  }
  if this.DealInfoStopped != nil {
    if err := XsDealInfoStoppedTypeFuncInit(this.DealInfoStopped, ctx); err != nil { return err }
  }
  if this.CreditorChanged != nil {
    if err := XsCreditorChangedTypeFuncInit(this.CreditorChanged, ctx); err != nil { return err }
  }
  if this.ServiceOrgChanged != nil {
    if err := XsServiceOrgChangedTypeFuncInit(this.ServiceOrgChanged, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsPersonEventType) Load(ctx *ParseContext) error {
  if err := XsPersonSubjectTypeFuncLoad(&this.Subject, ctx); err != nil { return err }
  if this.SubjectApplied != nil {
    if err := XsSubjectAppliedTypeFuncLoad(this.SubjectApplied, ctx); err != nil { return err }
  }
  if this.ApplicationChanged != nil {
    if err := XsApplicationChangedTypeFuncLoad(this.ApplicationChanged, ctx); err != nil { return err }
  }
  if this.ApplicationRejected != nil {
    if err := XsApplicationRejectedTypeFuncLoad(this.ApplicationRejected, ctx); err != nil { return err }
  }
  if this.MonetaryDealMade != nil {
    if err := XsMonetaryDealMadeTypeFuncLoad(this.MonetaryDealMade, ctx); err != nil { return err }
  }
  if this.SourceNonMonetaryDealMade != nil {
    if err := XsSourceNonMonetaryDealMadeTypeFuncLoad(this.SourceNonMonetaryDealMade, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealMade != nil {
    if err := XsNonMonetaryDealMadeTypeFuncLoad(this.NonMonetaryDealMade, ctx); err != nil { return err }
  }
  if this.LeasingDealMade != nil {
    if err := XsLeasingDealMadeTypeFuncLoad(this.LeasingDealMade, ctx); err != nil { return err }
  }
  if this.DebtClaimSubmitted != nil {
    if err := XsDebtClaimSubmittedTypeFuncLoad(this.DebtClaimSubmitted, ctx); err != nil { return err }
  }
  if this.DebtClaimChanged != nil {
    if err := XsDebtClaimChangedTypeFuncLoad(this.DebtClaimChanged, ctx); err != nil { return err }
  }
  if this.SubjectMainChanged != nil {
    if err := XsSubjectMainChangedTypeFuncLoad(this.SubjectMainChanged, ctx); err != nil { return err }
  }
  if this.SubjectSpecialChanged != nil {
    if err := XsSubjectSpecialChangedTypeFuncLoad(this.SubjectSpecialChanged, ctx); err != nil { return err }
  }
  if this.SubjectCapacityChanged != nil {
    if err := XsSubjectCapacityChangedTypeFuncLoad(this.SubjectCapacityChanged, ctx); err != nil { return err }
  }
  if this.SubjectBankruptcyChanged != nil {
    if err := XsSubjectBankruptcyChangedTypeFuncLoad(this.SubjectBankruptcyChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealChanged != nil {
    if err := XsMonetaryDealChangedTypeFuncLoad(this.MonetaryDealChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealChanged != nil {
    if err := XsNonMonetaryDealChangedTypeFuncLoad(this.NonMonetaryDealChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealPrincipalChanged != nil {
    if err := XsMonetaryDealPrincipalChangedTypeFuncLoad(this.MonetaryDealPrincipalChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealPrincipalChanged != nil {
    if err := XsNonMonetaryDealPrincipalChangedTypeFuncLoad(this.NonMonetaryDealPrincipalChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealPerformanceChanged != nil {
    if err := XsMonetaryDealPerformanceChangedTypeFuncLoad(this.MonetaryDealPerformanceChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealPerformanceChanged != nil {
    if err := XsNonMonetaryDealPerformanceChangedTypeFuncLoad(this.NonMonetaryDealPerformanceChanged, ctx); err != nil { return err }
  }
  if this.DealSecuringChanged != nil {
    if err := XsDealSecuringChangedTypeFuncLoad(this.DealSecuringChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealEnded != nil {
    if err := XsMonetaryDealEndedTypeFuncLoad(this.MonetaryDealEnded, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealEnded != nil {
    if err := XsNonMonetaryDealEndedTypeFuncLoad(this.NonMonetaryDealEnded, ctx); err != nil { return err }
  }
  if this.DealClaimChanged != nil {
    if err := XsDealClaimChangedTypeFuncLoad(this.DealClaimChanged, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyStarted != nil {
    if err := XsSourceBankruptcyStartedTypeFuncLoad(this.SourceBankruptcyStarted, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyChanged != nil {
    if err := XsSourceBankruptcyChangedTypeFuncLoad(this.SourceBankruptcyChanged, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyEnded != nil {
    if err := XsSourceBankruptcyEndedTypeFuncLoad(this.SourceBankruptcyEnded, ctx); err != nil { return err }
  }
  if this.SourceLiquidationStarted != nil {
    if err := XsSourceLiquidationStartedTypeFuncLoad(this.SourceLiquidationStarted, ctx); err != nil { return err }
  }
  if this.SourceLiquidationChanged != nil {
    if err := XsSourceLiquidationChangedTypeFuncLoad(this.SourceLiquidationChanged, ctx); err != nil { return err }
  }
  if this.SourceLiquidationEnded != nil {
    if err := XsSourceLiquidationEndedTypeFuncLoad(this.SourceLiquidationEnded, ctx); err != nil { return err }
  }
  if this.DealInfoStopped != nil {
    if err := XsDealInfoStoppedTypeFuncLoad(this.DealInfoStopped, ctx); err != nil { return err }
  }
  if this.CreditorChanged != nil {
    if err := XsCreditorChangedTypeFuncLoad(this.CreditorChanged, ctx); err != nil { return err }
  }
  if this.ServiceOrgChanged != nil {
    if err := XsServiceOrgChangedTypeFuncLoad(this.ServiceOrgChanged, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsPersonEventType) Validate(ctx *ParseContext) error {
  if err := XsPersonSubjectTypeFuncValidate(&this.Subject, ctx); err != nil { return err }
  if this.SubjectApplied != nil {
    if err := XsSubjectAppliedTypeFuncValidate(this.SubjectApplied, ctx); err != nil { return err }
  }
  if this.ApplicationChanged != nil {
    if err := XsApplicationChangedTypeFuncValidate(this.ApplicationChanged, ctx); err != nil { return err }
  }
  if this.ApplicationRejected != nil {
    if err := XsApplicationRejectedTypeFuncValidate(this.ApplicationRejected, ctx); err != nil { return err }
  }
  if this.MonetaryDealMade != nil {
    if err := XsMonetaryDealMadeTypeFuncValidate(this.MonetaryDealMade, ctx); err != nil { return err }
  }
  if this.SourceNonMonetaryDealMade != nil {
    if err := XsSourceNonMonetaryDealMadeTypeFuncValidate(this.SourceNonMonetaryDealMade, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealMade != nil {
    if err := XsNonMonetaryDealMadeTypeFuncValidate(this.NonMonetaryDealMade, ctx); err != nil { return err }
  }
  if this.LeasingDealMade != nil {
    if err := XsLeasingDealMadeTypeFuncValidate(this.LeasingDealMade, ctx); err != nil { return err }
  }
  if this.DebtClaimSubmitted != nil {
    if err := XsDebtClaimSubmittedTypeFuncValidate(this.DebtClaimSubmitted, ctx); err != nil { return err }
  }
  if this.DebtClaimChanged != nil {
    if err := XsDebtClaimChangedTypeFuncValidate(this.DebtClaimChanged, ctx); err != nil { return err }
  }
  if this.SubjectMainChanged != nil {
    if err := XsSubjectMainChangedTypeFuncValidate(this.SubjectMainChanged, ctx); err != nil { return err }
  }
  if this.SubjectSpecialChanged != nil {
    if err := XsSubjectSpecialChangedTypeFuncValidate(this.SubjectSpecialChanged, ctx); err != nil { return err }
  }
  if this.SubjectCapacityChanged != nil {
    if err := XsSubjectCapacityChangedTypeFuncValidate(this.SubjectCapacityChanged, ctx); err != nil { return err }
  }
  if this.SubjectBankruptcyChanged != nil {
    if err := XsSubjectBankruptcyChangedTypeFuncValidate(this.SubjectBankruptcyChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealChanged != nil {
    if err := XsMonetaryDealChangedTypeFuncValidate(this.MonetaryDealChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealChanged != nil {
    if err := XsNonMonetaryDealChangedTypeFuncValidate(this.NonMonetaryDealChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealPrincipalChanged != nil {
    if err := XsMonetaryDealPrincipalChangedTypeFuncValidate(this.MonetaryDealPrincipalChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealPrincipalChanged != nil {
    if err := XsNonMonetaryDealPrincipalChangedTypeFuncValidate(this.NonMonetaryDealPrincipalChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealPerformanceChanged != nil {
    if err := XsMonetaryDealPerformanceChangedTypeFuncValidate(this.MonetaryDealPerformanceChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealPerformanceChanged != nil {
    if err := XsNonMonetaryDealPerformanceChangedTypeFuncValidate(this.NonMonetaryDealPerformanceChanged, ctx); err != nil { return err }
  }
  if this.DealSecuringChanged != nil {
    if err := XsDealSecuringChangedTypeFuncValidate(this.DealSecuringChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealEnded != nil {
    if err := XsMonetaryDealEndedTypeFuncValidate(this.MonetaryDealEnded, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealEnded != nil {
    if err := XsNonMonetaryDealEndedTypeFuncValidate(this.NonMonetaryDealEnded, ctx); err != nil { return err }
  }
  if this.DealClaimChanged != nil {
    if err := XsDealClaimChangedTypeFuncValidate(this.DealClaimChanged, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyStarted != nil {
    if err := XsSourceBankruptcyStartedTypeFuncValidate(this.SourceBankruptcyStarted, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyChanged != nil {
    if err := XsSourceBankruptcyChangedTypeFuncValidate(this.SourceBankruptcyChanged, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyEnded != nil {
    if err := XsSourceBankruptcyEndedTypeFuncValidate(this.SourceBankruptcyEnded, ctx); err != nil { return err }
  }
  if this.SourceLiquidationStarted != nil {
    if err := XsSourceLiquidationStartedTypeFuncValidate(this.SourceLiquidationStarted, ctx); err != nil { return err }
  }
  if this.SourceLiquidationChanged != nil {
    if err := XsSourceLiquidationChangedTypeFuncValidate(this.SourceLiquidationChanged, ctx); err != nil { return err }
  }
  if this.SourceLiquidationEnded != nil {
    if err := XsSourceLiquidationEndedTypeFuncValidate(this.SourceLiquidationEnded, ctx); err != nil { return err }
  }
  if this.DealInfoStopped != nil {
    if err := XsDealInfoStoppedTypeFuncValidate(this.DealInfoStopped, ctx); err != nil { return err }
  }
  if this.CreditorChanged != nil {
    if err := XsCreditorChangedTypeFuncValidate(this.CreditorChanged, ctx); err != nil { return err }
  }
  if this.ServiceOrgChanged != nil {
    if err := XsServiceOrgChangedTypeFuncValidate(this.ServiceOrgChanged, ctx); err != nil { return err }
  }
  return nil
}

// Событие Кредитной Истории Юридического Лица
type XsOrgEventTypeFunc func(this *XsOrgEventType, ctx *ParseContext) error

var XsOrgEventTypeFuncInit XsOrgEventTypeFunc = (*XsOrgEventType).Init

var XsOrgEventTypeFuncLoad XsOrgEventTypeFunc = (*XsOrgEventType).Load

var XsOrgEventTypeFuncValidate XsOrgEventTypeFunc = (*XsOrgEventType).Validate

func (this *XsOrgEventType) Init(ctx *ParseContext) error {
  if err := XsOrgSubjectTypeFuncInit(&this.Subject, ctx); err != nil { return err }
  if this.SubjectApplied != nil {
    if err := XsOrgSubjectAppliedTypeFuncInit(this.SubjectApplied, ctx); err != nil { return err }
  }
  if this.ApplicationChanged != nil {
    if err := XsOrgApplicationChangedTypeFuncInit(this.ApplicationChanged, ctx); err != nil { return err }
  }
  if this.ApplicationRejected != nil {
    if err := XsOrgApplicationRejectedTypeFuncInit(this.ApplicationRejected, ctx); err != nil { return err }
  }
  if this.MonetaryDealMade != nil {
    if err := XsOrgMonetaryDealMadeTypeFuncInit(this.MonetaryDealMade, ctx); err != nil { return err }
  }
  if this.SourceNonMonetaryDealMade != nil {
    if err := XsOrgSourceNonMonetaryDealMadeTypeFuncInit(this.SourceNonMonetaryDealMade, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealMade != nil {
    if err := XsOrgNonMonetaryDealMadeTypeFuncInit(this.NonMonetaryDealMade, ctx); err != nil { return err }
  }
  if this.LeasingDealMade != nil {
    if err := XsOrgLeasingDealMadeTypeFuncInit(this.LeasingDealMade, ctx); err != nil { return err }
  }
  if this.DebtClaimSubmitted != nil {
    if err := XsOrgDebtClaimSubmittedTypeFuncInit(this.DebtClaimSubmitted, ctx); err != nil { return err }
  }
  if this.DebtClaimChanged != nil {
    if err := XsOrgDebtClaimChangedTypeFuncInit(this.DebtClaimChanged, ctx); err != nil { return err }
  }
  if this.SubjectMainChanged != nil {
    if err := XsOrgSubjectMainChangedTypeFuncInit(this.SubjectMainChanged, ctx); err != nil { return err }
  }
  if this.SubjectSpecialChanged != nil {
    if err := XsOrgSubjectSpecialChangedTypeFuncInit(this.SubjectSpecialChanged, ctx); err != nil { return err }
  }
  if this.SubjectBankruptcyChanged != nil {
    if err := XsOrgSubjectBankruptcyChangedTypeFuncInit(this.SubjectBankruptcyChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealChanged != nil {
    if err := XsOrgMonetaryDealChangedTypeFuncInit(this.MonetaryDealChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealChanged != nil {
    if err := XsOrgNonMonetaryDealChangedTypeFuncInit(this.NonMonetaryDealChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealPrincipalChanged != nil {
    if err := XsOrgMonetaryDealPrincipalChangedTypeFuncInit(this.MonetaryDealPrincipalChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealPrincipalChanged != nil {
    if err := XsOrgNonMonetaryDealPrincipalChangedTypeFuncInit(this.NonMonetaryDealPrincipalChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealPerformanceChanged != nil {
    if err := XsOrgMonetaryDealParticipationChangedTypeFuncInit(this.MonetaryDealPerformanceChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealPerformanceChanged != nil {
    if err := XsOrgNonMonetaryDealParticipationChangedTypeFuncInit(this.NonMonetaryDealPerformanceChanged, ctx); err != nil { return err }
  }
  if this.DealSecuringChanged != nil {
    if err := XsOrgDealSecuringChangedTypeFuncInit(this.DealSecuringChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealEnded != nil {
    if err := XsOrgMonetaryDealEndedTypeFuncInit(this.MonetaryDealEnded, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealEnded != nil {
    if err := XsOrgNonMonetaryDealEndedTypeFuncInit(this.NonMonetaryDealEnded, ctx); err != nil { return err }
  }
  if this.DealClaimChanged != nil {
    if err := XsOrgDealClaimChangedTypeFuncInit(this.DealClaimChanged, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyStarted != nil {
    if err := XsOrgSourceBankruptcyStartedTypeFuncInit(this.SourceBankruptcyStarted, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyChanged != nil {
    if err := XsOrgSourceBankruptcyChangedTypeFuncInit(this.SourceBankruptcyChanged, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyEnded != nil {
    if err := XsOrgSourceBankruptcyEndedTypeFuncInit(this.SourceBankruptcyEnded, ctx); err != nil { return err }
  }
  if this.SourceLiquidationStarted != nil {
    if err := XsOrgSourceLiquidationStartedTypeFuncInit(this.SourceLiquidationStarted, ctx); err != nil { return err }
  }
  if this.SourceLiquidationChanged != nil {
    if err := XsOrgSourceLiquidationChangedTypeFuncInit(this.SourceLiquidationChanged, ctx); err != nil { return err }
  }
  if this.SourceLiquidationEnded != nil {
    if err := XsOrgSourceLiquidationEndedTypeFuncInit(this.SourceLiquidationEnded, ctx); err != nil { return err }
  }
  if this.DealInfoStopped != nil {
    if err := XsOrgDealInfoStoppedTypeFuncInit(this.DealInfoStopped, ctx); err != nil { return err }
  }
  if this.CreditorChanged != nil {
    if err := XsOrgCreditorChangedTypeFuncInit(this.CreditorChanged, ctx); err != nil { return err }
  }
  if this.ServiceOrgChanged != nil {
    if err := XsOrgServiceOrgChangedTypeFuncInit(this.ServiceOrgChanged, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsOrgEventType) Load(ctx *ParseContext) error {
  if err := XsOrgSubjectTypeFuncLoad(&this.Subject, ctx); err != nil { return err }
  if this.SubjectApplied != nil {
    if err := XsOrgSubjectAppliedTypeFuncLoad(this.SubjectApplied, ctx); err != nil { return err }
  }
  if this.ApplicationChanged != nil {
    if err := XsOrgApplicationChangedTypeFuncLoad(this.ApplicationChanged, ctx); err != nil { return err }
  }
  if this.ApplicationRejected != nil {
    if err := XsOrgApplicationRejectedTypeFuncLoad(this.ApplicationRejected, ctx); err != nil { return err }
  }
  if this.MonetaryDealMade != nil {
    if err := XsOrgMonetaryDealMadeTypeFuncLoad(this.MonetaryDealMade, ctx); err != nil { return err }
  }
  if this.SourceNonMonetaryDealMade != nil {
    if err := XsOrgSourceNonMonetaryDealMadeTypeFuncLoad(this.SourceNonMonetaryDealMade, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealMade != nil {
    if err := XsOrgNonMonetaryDealMadeTypeFuncLoad(this.NonMonetaryDealMade, ctx); err != nil { return err }
  }
  if this.LeasingDealMade != nil {
    if err := XsOrgLeasingDealMadeTypeFuncLoad(this.LeasingDealMade, ctx); err != nil { return err }
  }
  if this.DebtClaimSubmitted != nil {
    if err := XsOrgDebtClaimSubmittedTypeFuncLoad(this.DebtClaimSubmitted, ctx); err != nil { return err }
  }
  if this.DebtClaimChanged != nil {
    if err := XsOrgDebtClaimChangedTypeFuncLoad(this.DebtClaimChanged, ctx); err != nil { return err }
  }
  if this.SubjectMainChanged != nil {
    if err := XsOrgSubjectMainChangedTypeFuncLoad(this.SubjectMainChanged, ctx); err != nil { return err }
  }
  if this.SubjectSpecialChanged != nil {
    if err := XsOrgSubjectSpecialChangedTypeFuncLoad(this.SubjectSpecialChanged, ctx); err != nil { return err }
  }
  if this.SubjectBankruptcyChanged != nil {
    if err := XsOrgSubjectBankruptcyChangedTypeFuncLoad(this.SubjectBankruptcyChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealChanged != nil {
    if err := XsOrgMonetaryDealChangedTypeFuncLoad(this.MonetaryDealChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealChanged != nil {
    if err := XsOrgNonMonetaryDealChangedTypeFuncLoad(this.NonMonetaryDealChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealPrincipalChanged != nil {
    if err := XsOrgMonetaryDealPrincipalChangedTypeFuncLoad(this.MonetaryDealPrincipalChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealPrincipalChanged != nil {
    if err := XsOrgNonMonetaryDealPrincipalChangedTypeFuncLoad(this.NonMonetaryDealPrincipalChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealPerformanceChanged != nil {
    if err := XsOrgMonetaryDealParticipationChangedTypeFuncLoad(this.MonetaryDealPerformanceChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealPerformanceChanged != nil {
    if err := XsOrgNonMonetaryDealParticipationChangedTypeFuncLoad(this.NonMonetaryDealPerformanceChanged, ctx); err != nil { return err }
  }
  if this.DealSecuringChanged != nil {
    if err := XsOrgDealSecuringChangedTypeFuncLoad(this.DealSecuringChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealEnded != nil {
    if err := XsOrgMonetaryDealEndedTypeFuncLoad(this.MonetaryDealEnded, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealEnded != nil {
    if err := XsOrgNonMonetaryDealEndedTypeFuncLoad(this.NonMonetaryDealEnded, ctx); err != nil { return err }
  }
  if this.DealClaimChanged != nil {
    if err := XsOrgDealClaimChangedTypeFuncLoad(this.DealClaimChanged, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyStarted != nil {
    if err := XsOrgSourceBankruptcyStartedTypeFuncLoad(this.SourceBankruptcyStarted, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyChanged != nil {
    if err := XsOrgSourceBankruptcyChangedTypeFuncLoad(this.SourceBankruptcyChanged, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyEnded != nil {
    if err := XsOrgSourceBankruptcyEndedTypeFuncLoad(this.SourceBankruptcyEnded, ctx); err != nil { return err }
  }
  if this.SourceLiquidationStarted != nil {
    if err := XsOrgSourceLiquidationStartedTypeFuncLoad(this.SourceLiquidationStarted, ctx); err != nil { return err }
  }
  if this.SourceLiquidationChanged != nil {
    if err := XsOrgSourceLiquidationChangedTypeFuncLoad(this.SourceLiquidationChanged, ctx); err != nil { return err }
  }
  if this.SourceLiquidationEnded != nil {
    if err := XsOrgSourceLiquidationEndedTypeFuncLoad(this.SourceLiquidationEnded, ctx); err != nil { return err }
  }
  if this.DealInfoStopped != nil {
    if err := XsOrgDealInfoStoppedTypeFuncLoad(this.DealInfoStopped, ctx); err != nil { return err }
  }
  if this.CreditorChanged != nil {
    if err := XsOrgCreditorChangedTypeFuncLoad(this.CreditorChanged, ctx); err != nil { return err }
  }
  if this.ServiceOrgChanged != nil {
    if err := XsOrgServiceOrgChangedTypeFuncLoad(this.ServiceOrgChanged, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsOrgEventType) Validate(ctx *ParseContext) error {
  if err := XsOrgSubjectTypeFuncValidate(&this.Subject, ctx); err != nil { return err }
  if this.SubjectApplied != nil {
    if err := XsOrgSubjectAppliedTypeFuncValidate(this.SubjectApplied, ctx); err != nil { return err }
  }
  if this.ApplicationChanged != nil {
    if err := XsOrgApplicationChangedTypeFuncValidate(this.ApplicationChanged, ctx); err != nil { return err }
  }
  if this.ApplicationRejected != nil {
    if err := XsOrgApplicationRejectedTypeFuncValidate(this.ApplicationRejected, ctx); err != nil { return err }
  }
  if this.MonetaryDealMade != nil {
    if err := XsOrgMonetaryDealMadeTypeFuncValidate(this.MonetaryDealMade, ctx); err != nil { return err }
  }
  if this.SourceNonMonetaryDealMade != nil {
    if err := XsOrgSourceNonMonetaryDealMadeTypeFuncValidate(this.SourceNonMonetaryDealMade, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealMade != nil {
    if err := XsOrgNonMonetaryDealMadeTypeFuncValidate(this.NonMonetaryDealMade, ctx); err != nil { return err }
  }
  if this.LeasingDealMade != nil {
    if err := XsOrgLeasingDealMadeTypeFuncValidate(this.LeasingDealMade, ctx); err != nil { return err }
  }
  if this.DebtClaimSubmitted != nil {
    if err := XsOrgDebtClaimSubmittedTypeFuncValidate(this.DebtClaimSubmitted, ctx); err != nil { return err }
  }
  if this.DebtClaimChanged != nil {
    if err := XsOrgDebtClaimChangedTypeFuncValidate(this.DebtClaimChanged, ctx); err != nil { return err }
  }
  if this.SubjectMainChanged != nil {
    if err := XsOrgSubjectMainChangedTypeFuncValidate(this.SubjectMainChanged, ctx); err != nil { return err }
  }
  if this.SubjectSpecialChanged != nil {
    if err := XsOrgSubjectSpecialChangedTypeFuncValidate(this.SubjectSpecialChanged, ctx); err != nil { return err }
  }
  if this.SubjectBankruptcyChanged != nil {
    if err := XsOrgSubjectBankruptcyChangedTypeFuncValidate(this.SubjectBankruptcyChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealChanged != nil {
    if err := XsOrgMonetaryDealChangedTypeFuncValidate(this.MonetaryDealChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealChanged != nil {
    if err := XsOrgNonMonetaryDealChangedTypeFuncValidate(this.NonMonetaryDealChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealPrincipalChanged != nil {
    if err := XsOrgMonetaryDealPrincipalChangedTypeFuncValidate(this.MonetaryDealPrincipalChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealPrincipalChanged != nil {
    if err := XsOrgNonMonetaryDealPrincipalChangedTypeFuncValidate(this.NonMonetaryDealPrincipalChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealPerformanceChanged != nil {
    if err := XsOrgMonetaryDealParticipationChangedTypeFuncValidate(this.MonetaryDealPerformanceChanged, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealPerformanceChanged != nil {
    if err := XsOrgNonMonetaryDealParticipationChangedTypeFuncValidate(this.NonMonetaryDealPerformanceChanged, ctx); err != nil { return err }
  }
  if this.DealSecuringChanged != nil {
    if err := XsOrgDealSecuringChangedTypeFuncValidate(this.DealSecuringChanged, ctx); err != nil { return err }
  }
  if this.MonetaryDealEnded != nil {
    if err := XsOrgMonetaryDealEndedTypeFuncValidate(this.MonetaryDealEnded, ctx); err != nil { return err }
  }
  if this.NonMonetaryDealEnded != nil {
    if err := XsOrgNonMonetaryDealEndedTypeFuncValidate(this.NonMonetaryDealEnded, ctx); err != nil { return err }
  }
  if this.DealClaimChanged != nil {
    if err := XsOrgDealClaimChangedTypeFuncValidate(this.DealClaimChanged, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyStarted != nil {
    if err := XsOrgSourceBankruptcyStartedTypeFuncValidate(this.SourceBankruptcyStarted, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyChanged != nil {
    if err := XsOrgSourceBankruptcyChangedTypeFuncValidate(this.SourceBankruptcyChanged, ctx); err != nil { return err }
  }
  if this.SourceBankruptcyEnded != nil {
    if err := XsOrgSourceBankruptcyEndedTypeFuncValidate(this.SourceBankruptcyEnded, ctx); err != nil { return err }
  }
  if this.SourceLiquidationStarted != nil {
    if err := XsOrgSourceLiquidationStartedTypeFuncValidate(this.SourceLiquidationStarted, ctx); err != nil { return err }
  }
  if this.SourceLiquidationChanged != nil {
    if err := XsOrgSourceLiquidationChangedTypeFuncValidate(this.SourceLiquidationChanged, ctx); err != nil { return err }
  }
  if this.SourceLiquidationEnded != nil {
    if err := XsOrgSourceLiquidationEndedTypeFuncValidate(this.SourceLiquidationEnded, ctx); err != nil { return err }
  }
  if this.DealInfoStopped != nil {
    if err := XsOrgDealInfoStoppedTypeFuncValidate(this.DealInfoStopped, ctx); err != nil { return err }
  }
  if this.CreditorChanged != nil {
    if err := XsOrgCreditorChangedTypeFuncValidate(this.CreditorChanged, ctx); err != nil { return err }
  }
  if this.ServiceOrgChanged != nil {
    if err := XsOrgServiceOrgChangedTypeFuncValidate(this.ServiceOrgChanged, ctx); err != nil { return err }
  }
  return nil
}

// Сведения об источнике формирования кредитной истории. infoDate и один из блоков 46, 47, 48 (orgSource, personSource, commissionerSource)
type XsSourceTypeFunc func(this *XsSourceType, ctx *ParseContext) error

var XsSourceTypeFuncInit XsSourceTypeFunc = (*XsSourceType).Init

var XsSourceTypeFuncLoad XsSourceTypeFunc = (*XsSourceType).Load

var XsSourceTypeFuncValidate XsSourceTypeFunc = (*XsSourceType).Validate

func (this *XsSourceType) Init(ctx *ParseContext) error {
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncInit(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncInit(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncInit(this.CommissionerSource, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsSourceType) Load(ctx *ParseContext) error {
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncLoad(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncLoad(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncLoad(this.CommissionerSource, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsSourceType) Validate(ctx *ParseContext) error {
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncValidate(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncValidate(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncValidate(this.CommissionerSource, ctx); err != nil { return err }
  }
  return nil
}

// 1.1 Субъект обратился к источнику с предложением совершить сделку
type XsSubjectAppliedTypeFunc func(this *XsSubjectAppliedType, ctx *ParseContext) error

var XsSubjectAppliedTypeFuncInit XsSubjectAppliedTypeFunc = (*XsSubjectAppliedType).Init

var XsSubjectAppliedTypeFuncLoad XsSubjectAppliedTypeFunc = (*XsSubjectAppliedType).Load

var XsSubjectAppliedTypeFuncValidate XsSubjectAppliedTypeFunc = (*XsSubjectAppliedType).Validate

func (this *XsSubjectAppliedType) Init(ctx *ParseContext) error {
  if err := XsApplicationTypeFuncInit(&this.Application, ctx); err != nil { return err }
  return nil
}

func (this *XsSubjectAppliedType) Load(ctx *ParseContext) error {
  foundApplication := -1
  for i, v := range ctx.person.Application {
    if this.Application.Same(&v) {
      ctx.person.Application[i] = this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.person.Application = append(ctx.person.Application, this.Application)
  }
  if ctx.op == "D" && foundApplication != -1 {
    ctx.person.Application = append(ctx.person.Application[:foundApplication], ctx.person.Application[foundApplication+1:]...)
  }
  return nil
}

func (this *XsSubjectAppliedType) Validate(ctx *ParseContext) error {
  if err := XsApplicationTypeFuncValidate(&this.Application, ctx); err != nil { return err }
  return nil
}

// 1.2 Источник одобрил обращение субъекта (направил ему оферту) или изменились сведения об обращении
type XsApplicationChangedTypeFunc func(this *XsApplicationChangedType, ctx *ParseContext) error

var XsApplicationChangedTypeFuncInit XsApplicationChangedTypeFunc = (*XsApplicationChangedType).Init

var XsApplicationChangedTypeFuncLoad XsApplicationChangedTypeFunc = (*XsApplicationChangedType).Load

var XsApplicationChangedTypeFuncValidate XsApplicationChangedTypeFunc = (*XsApplicationChangedType).Validate

func (this *XsApplicationChangedType) Init(ctx *ParseContext) error {
  if err := XsApplicationTypeFuncInit(&this.Application, ctx); err != nil { return err }
  return nil
}

func (this *XsApplicationChangedType) Load(ctx *ParseContext) error {
  foundApplication := -1
  for i, v := range ctx.person.Application {
    if this.Application.Same(&v) {
      ctx.person.Application[i] = this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.person.Application = append(ctx.person.Application, this.Application)
  }
  return nil
}

func (this *XsApplicationChangedType) Validate(ctx *ParseContext) error {
  if err := XsApplicationTypeFuncValidate(&this.Application, ctx); err != nil { return err }
  return nil
}

// 1.3 Источник отказался от совершения сделки по обращению субъекта
type XsApplicationRejectedTypeFunc func(this *XsApplicationRejectedType, ctx *ParseContext) error

var XsApplicationRejectedTypeFuncInit XsApplicationRejectedTypeFunc = (*XsApplicationRejectedType).Init

var XsApplicationRejectedTypeFuncLoad XsApplicationRejectedTypeFunc = (*XsApplicationRejectedType).Load

var XsApplicationRejectedTypeFuncValidate XsApplicationRejectedTypeFunc = (*XsApplicationRejectedType).Validate

func (this *XsApplicationRejectedType) Init(ctx *ParseContext) error {
  if err := XsApplicationTypeFuncInit(&this.Application, ctx); err != nil { return err }
  if err := XsAppRejectTypeFuncInit(&this.AppReject, ctx); err != nil { return err }
  return nil
}

func (this *XsApplicationRejectedType) Load(ctx *ParseContext) error {
  foundApplication := -1
  for i, v := range ctx.person.Application {
    if this.Application.Same(&v) {
      ctx.person.Application[i] = this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.person.Application = append(ctx.person.Application, this.Application)
  }
  foundAppReject := -1
  for i, v := range ctx.person.AppReject {
    if this.AppReject.Same(&v) {
      ctx.person.AppReject[i] = this.AppReject
      foundAppReject = i
    }
  }
  if ctx.op != "D" && foundAppReject == -1 {
    ctx.person.AppReject = append(ctx.person.AppReject, this.AppReject)
  }
  if ctx.op == "D" && foundAppReject != -1 {
    ctx.person.AppReject = append(ctx.person.AppReject[:foundAppReject], ctx.person.AppReject[foundAppReject+1:]...)
  }
  return nil
}

func (this *XsApplicationRejectedType) Validate(ctx *ParseContext) error {
  if err := XsApplicationTypeFuncValidate(&this.Application, ctx); err != nil { return err }
  if err := XsAppRejectTypeFuncValidate(&this.AppReject, ctx); err != nil { return err }
  return nil
}

// 1.4 Субъект и источник совершили сделку, кроме договора лизинга и поручительства по лизингу для денежного обязательства субъекта
type XsMonetaryDealMadeTypeFunc func(this *XsMonetaryDealMadeType, ctx *ParseContext) error

var XsMonetaryDealMadeTypeFuncInit XsMonetaryDealMadeTypeFunc = (*XsMonetaryDealMadeType).Init

var XsMonetaryDealMadeTypeFuncLoad XsMonetaryDealMadeTypeFunc = (*XsMonetaryDealMadeType).Load

var XsMonetaryDealMadeTypeFuncValidate XsMonetaryDealMadeTypeFunc = (*XsMonetaryDealMadeType).Validate

func (this *XsMonetaryDealMadeType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsAddressTypeFuncInit(&this.Address, ctx); err != nil { return err }
  if err := XsFactAddressTypeFuncInit(&this.FactAddress, ctx); err != nil { return err }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsSoleProprietorTypeFuncInit(&this.SoleProprietor, ctx); err != nil { return err }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncInit(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncInit(this.Application, ctx); err != nil { return err }
  }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncInit(&this.PaymentTerms, ctx); err != nil { return err }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncInit(this.TotalCost, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsMonetaryDealMadeType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Address.Same(&ctx.person.Address) {
    ctx.person.Address = this.Address
  }
  if   !this.FactAddress.Same(ctx.person.FactAddress) {
    ctx.person.FactAddress = &this.FactAddress
  }
  foundContacts := make(map[int]bool)
  for i, v := range ctx.person.Contact {
    for k, w := range this.Contact {
      if w.Same(&v) {
        ctx.person.Contact[i] = w
        foundContacts[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Contact {
      if !foundContacts[i] {
        ctx.person.Contact = append(ctx.person.Contact, w)
      }
    }
  }
  if   !this.SoleProprietor.Same(ctx.person.SoleProprietor) {
    ctx.person.SoleProprietor = &this.SoleProprietor
  }
  foundLegalCapacitys := make(map[int]bool)
  for i, v := range ctx.person.LegalCapacity {
    for k, w := range this.LegalCapacity {
      if w.Same(&v) {
        ctx.person.LegalCapacity[i] = w
        foundLegalCapacitys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.LegalCapacity {
      if !foundLegalCapacitys[i] {
        ctx.person.LegalCapacity = append(ctx.person.LegalCapacity, w)
      }
    }
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if   !this.JointDebtors.Same(ctx.personDeal.JointDebtors) {
    ctx.personDeal.JointDebtors = &this.JointDebtors
  }
  if   !this.Accounting.Same(ctx.personDeal.Accounting) {
    ctx.personDeal.Accounting = &this.Accounting
  }
  foundApplication := -1
  for i, v := range ctx.person.Application {
    if this.Application.Same(&v) {
      ctx.person.Application[i] = *this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.person.Application = append(ctx.person.Application, *this.Application)
  }
  if   !this.Amount.Same(ctx.personDeal.Amount) {
    ctx.personDeal.Amount = &this.Amount
  }
  if   !this.PaymentTerms.Same(ctx.personDeal.PaymentTerms) {
    ctx.personDeal.PaymentTerms = &this.PaymentTerms
  }
  if   !this.TotalCost.Same(ctx.personDeal.TotalCost) {
    ctx.personDeal.TotalCost = this.TotalCost
  }
  return nil
}

func (this *XsMonetaryDealMadeType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsAddressTypeFuncValidate(&this.Address, ctx); err != nil { return err }
  if err := XsFactAddressTypeFuncValidate(&this.FactAddress, ctx); err != nil { return err }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsSoleProprietorTypeFuncValidate(&this.SoleProprietor, ctx); err != nil { return err }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncValidate(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncValidate(this.Application, ctx); err != nil { return err }
  }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncValidate(&this.PaymentTerms, ctx); err != nil { return err }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncValidate(this.TotalCost, ctx); err != nil { return err }
  }
  return nil
}

// 1.4 Субъект и источник совершили сделку, кроме договора лизинга и поручительства по лизингу для неденежного обязательства источника
type XsSourceNonMonetaryDealMadeTypeFunc func(this *XsSourceNonMonetaryDealMadeType, ctx *ParseContext) error

var XsSourceNonMonetaryDealMadeTypeFuncInit XsSourceNonMonetaryDealMadeTypeFunc = (*XsSourceNonMonetaryDealMadeType).Init

var XsSourceNonMonetaryDealMadeTypeFuncLoad XsSourceNonMonetaryDealMadeTypeFunc = (*XsSourceNonMonetaryDealMadeType).Load

var XsSourceNonMonetaryDealMadeTypeFuncValidate XsSourceNonMonetaryDealMadeTypeFunc = (*XsSourceNonMonetaryDealMadeType).Validate

func (this *XsSourceNonMonetaryDealMadeType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsAddressTypeFuncInit(&this.Address, ctx); err != nil { return err }
  if err := XsFactAddressTypeFuncInit(&this.FactAddress, ctx); err != nil { return err }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsSoleProprietorTypeFuncInit(&this.SoleProprietor, ctx); err != nil { return err }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncInit(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncInit(this.Application, ctx); err != nil { return err }
  }
  if err := XsSourceNonMonetaryObligationTypeFuncInit(&this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

func (this *XsSourceNonMonetaryDealMadeType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Address.Same(&ctx.person.Address) {
    ctx.person.Address = this.Address
  }
  if   !this.FactAddress.Same(ctx.person.FactAddress) {
    ctx.person.FactAddress = &this.FactAddress
  }
  foundContacts := make(map[int]bool)
  for i, v := range ctx.person.Contact {
    for k, w := range this.Contact {
      if w.Same(&v) {
        ctx.person.Contact[i] = w
        foundContacts[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Contact {
      if !foundContacts[i] {
        ctx.person.Contact = append(ctx.person.Contact, w)
      }
    }
  }
  if   !this.SoleProprietor.Same(ctx.person.SoleProprietor) {
    ctx.person.SoleProprietor = &this.SoleProprietor
  }
  foundLegalCapacitys := make(map[int]bool)
  for i, v := range ctx.person.LegalCapacity {
    for k, w := range this.LegalCapacity {
      if w.Same(&v) {
        ctx.person.LegalCapacity[i] = w
        foundLegalCapacitys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.LegalCapacity {
      if !foundLegalCapacitys[i] {
        ctx.person.LegalCapacity = append(ctx.person.LegalCapacity, w)
      }
    }
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if   !this.JointDebtors.Same(ctx.personDeal.JointDebtors) {
    ctx.personDeal.JointDebtors = &this.JointDebtors
  }
  if   !this.Accounting.Same(ctx.personDeal.Accounting) {
    ctx.personDeal.Accounting = &this.Accounting
  }
  foundApplication := -1
  for i, v := range ctx.person.Application {
    if this.Application.Same(&v) {
      ctx.person.Application[i] = *this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.person.Application = append(ctx.person.Application, *this.Application)
  }
  if   !this.SourceNonMonetaryObligation.Same(ctx.personDeal.SourceNonMonetaryObligation) {
    ctx.personDeal.SourceNonMonetaryObligation = &this.SourceNonMonetaryObligation
  }
  return nil
}

func (this *XsSourceNonMonetaryDealMadeType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsAddressTypeFuncValidate(&this.Address, ctx); err != nil { return err }
  if err := XsFactAddressTypeFuncValidate(&this.FactAddress, ctx); err != nil { return err }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsSoleProprietorTypeFuncValidate(&this.SoleProprietor, ctx); err != nil { return err }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncValidate(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncValidate(this.Application, ctx); err != nil { return err }
  }
  if err := XsSourceNonMonetaryObligationTypeFuncValidate(&this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

// 1.4 Субъект и источник совершили сделку, кроме договора лизинга и поручительства по лизингу для неденежного обязательства субъекта  
type XsNonMonetaryDealMadeTypeFunc func(this *XsNonMonetaryDealMadeType, ctx *ParseContext) error

var XsNonMonetaryDealMadeTypeFuncInit XsNonMonetaryDealMadeTypeFunc = (*XsNonMonetaryDealMadeType).Init

var XsNonMonetaryDealMadeTypeFuncLoad XsNonMonetaryDealMadeTypeFunc = (*XsNonMonetaryDealMadeType).Load

var XsNonMonetaryDealMadeTypeFuncValidate XsNonMonetaryDealMadeTypeFunc = (*XsNonMonetaryDealMadeType).Validate

func (this *XsNonMonetaryDealMadeType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsAddressTypeFuncInit(&this.Address, ctx); err != nil { return err }
  if err := XsFactAddressTypeFuncInit(&this.FactAddress, ctx); err != nil { return err }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsSoleProprietorTypeFuncInit(&this.SoleProprietor, ctx); err != nil { return err }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncInit(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncInit(this.Application, ctx); err != nil { return err }
  }
  if err := XsSubjectNonMonetaryObligationTypeFuncInit(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

func (this *XsNonMonetaryDealMadeType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Address.Same(&ctx.person.Address) {
    ctx.person.Address = this.Address
  }
  if   !this.FactAddress.Same(ctx.person.FactAddress) {
    ctx.person.FactAddress = &this.FactAddress
  }
  foundContacts := make(map[int]bool)
  for i, v := range ctx.person.Contact {
    for k, w := range this.Contact {
      if w.Same(&v) {
        ctx.person.Contact[i] = w
        foundContacts[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Contact {
      if !foundContacts[i] {
        ctx.person.Contact = append(ctx.person.Contact, w)
      }
    }
  }
  if   !this.SoleProprietor.Same(ctx.person.SoleProprietor) {
    ctx.person.SoleProprietor = &this.SoleProprietor
  }
  foundLegalCapacitys := make(map[int]bool)
  for i, v := range ctx.person.LegalCapacity {
    for k, w := range this.LegalCapacity {
      if w.Same(&v) {
        ctx.person.LegalCapacity[i] = w
        foundLegalCapacitys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.LegalCapacity {
      if !foundLegalCapacitys[i] {
        ctx.person.LegalCapacity = append(ctx.person.LegalCapacity, w)
      }
    }
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if   !this.JointDebtors.Same(ctx.personDeal.JointDebtors) {
    ctx.personDeal.JointDebtors = &this.JointDebtors
  }
  if   !this.Accounting.Same(ctx.personDeal.Accounting) {
    ctx.personDeal.Accounting = &this.Accounting
  }
  foundApplication := -1
  for i, v := range ctx.person.Application {
    if this.Application.Same(&v) {
      ctx.person.Application[i] = *this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.person.Application = append(ctx.person.Application, *this.Application)
  }
  if   !this.SubjectNonMonetaryObligation.Same(ctx.personDeal.SubjectNonMonetaryObligation) {
    ctx.personDeal.SubjectNonMonetaryObligation = &this.SubjectNonMonetaryObligation
  }
  return nil
}

func (this *XsNonMonetaryDealMadeType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsAddressTypeFuncValidate(&this.Address, ctx); err != nil { return err }
  if err := XsFactAddressTypeFuncValidate(&this.FactAddress, ctx); err != nil { return err }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsSoleProprietorTypeFuncValidate(&this.SoleProprietor, ctx); err != nil { return err }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncValidate(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncValidate(this.Application, ctx); err != nil { return err }
  }
  if err := XsSubjectNonMonetaryObligationTypeFuncValidate(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

// 1.4.1 Субъект и источник заключили договор лизинга либо поручительства по лизингу и предмет лизинга передан лизингополучателю
type XsLeasingDealMadeTypeFunc func(this *XsLeasingDealMadeType, ctx *ParseContext) error

var XsLeasingDealMadeTypeFuncInit XsLeasingDealMadeTypeFunc = (*XsLeasingDealMadeType).Init

var XsLeasingDealMadeTypeFuncLoad XsLeasingDealMadeTypeFunc = (*XsLeasingDealMadeType).Load

var XsLeasingDealMadeTypeFuncValidate XsLeasingDealMadeTypeFunc = (*XsLeasingDealMadeType).Validate

func (this *XsLeasingDealMadeType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsAddressTypeFuncInit(&this.Address, ctx); err != nil { return err }
  if err := XsFactAddressTypeFuncInit(&this.FactAddress, ctx); err != nil { return err }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsSoleProprietorTypeFuncInit(&this.SoleProprietor, ctx); err != nil { return err }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncInit(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncInit(this.Application, ctx); err != nil { return err }
  }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncInit(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncInit(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncInit(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncInit(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncInit(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsSourceNonMonetaryObligationTypeFuncInit(&this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsLeasingDealMadeType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Address.Same(&ctx.person.Address) {
    ctx.person.Address = this.Address
  }
  if   !this.FactAddress.Same(ctx.person.FactAddress) {
    ctx.person.FactAddress = &this.FactAddress
  }
  foundContacts := make(map[int]bool)
  for i, v := range ctx.person.Contact {
    for k, w := range this.Contact {
      if w.Same(&v) {
        ctx.person.Contact[i] = w
        foundContacts[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Contact {
      if !foundContacts[i] {
        ctx.person.Contact = append(ctx.person.Contact, w)
      }
    }
  }
  if   !this.SoleProprietor.Same(ctx.person.SoleProprietor) {
    ctx.person.SoleProprietor = &this.SoleProprietor
  }
  foundLegalCapacitys := make(map[int]bool)
  for i, v := range ctx.person.LegalCapacity {
    for k, w := range this.LegalCapacity {
      if w.Same(&v) {
        ctx.person.LegalCapacity[i] = w
        foundLegalCapacitys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.LegalCapacity {
      if !foundLegalCapacitys[i] {
        ctx.person.LegalCapacity = append(ctx.person.LegalCapacity, w)
      }
    }
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if   !this.JointDebtors.Same(ctx.personDeal.JointDebtors) {
    ctx.personDeal.JointDebtors = &this.JointDebtors
  }
  if   !this.Accounting.Same(ctx.personDeal.Accounting) {
    ctx.personDeal.Accounting = &this.Accounting
  }
  foundApplication := -1
  for i, v := range ctx.person.Application {
    if this.Application.Same(&v) {
      ctx.person.Application[i] = *this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.person.Application = append(ctx.person.Application, *this.Application)
  }
  if   !this.Amount.Same(ctx.personDeal.Amount) {
    ctx.personDeal.Amount = &this.Amount
  }
  if   !this.PaymentTerms.Same(ctx.personDeal.PaymentTerms) {
    ctx.personDeal.PaymentTerms = &this.PaymentTerms
  }
  if   !this.FundDate.Same(ctx.personDeal.FundDate) {
    ctx.personDeal.FundDate = &this.FundDate
  }
  foundArrear := -1
  for i, v := range ctx.personDeal.Arrear {
    if this.Arrear.Same(&v) {
      ctx.personDeal.Arrear[i] = this.Arrear
      foundArrear = i
    }
  }
  if ctx.op != "D" && foundArrear == -1 {
    ctx.personDeal.Arrear = append(ctx.personDeal.Arrear, this.Arrear)
  }
  foundDueArrear := -1
  for i, v := range ctx.personDeal.DueArrear {
    if this.DueArrear.Same(&v) {
      ctx.personDeal.DueArrear[i] = this.DueArrear
      foundDueArrear = i
    }
  }
  if ctx.op != "D" && foundDueArrear == -1 {
    ctx.personDeal.DueArrear = append(ctx.personDeal.DueArrear, this.DueArrear)
  }
  foundPastdueArrear := -1
  for i, v := range ctx.personDeal.PastdueArrear {
    if this.PastdueArrear.Same(&v) {
      ctx.personDeal.PastdueArrear[i] = this.PastdueArrear
      foundPastdueArrear = i
    }
  }
  if ctx.op != "D" && foundPastdueArrear == -1 {
    ctx.personDeal.PastdueArrear = append(ctx.personDeal.PastdueArrear, this.PastdueArrear)
  }
  foundPaymentss := make(map[int]bool)
  for i, v := range ctx.personDeal.Payments {
    for k, w := range this.Payments {
      if w.Same(&v) {
        ctx.personDeal.Payments[i] = w
        foundPaymentss[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Payments {
      if !foundPaymentss[i] {
        ctx.personDeal.Payments = append(ctx.personDeal.Payments, w)
      }
    }
  }
  if   !this.SourceNonMonetaryObligation.Same(ctx.personDeal.SourceNonMonetaryObligation) {
    ctx.personDeal.SourceNonMonetaryObligation = &this.SourceNonMonetaryObligation
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  return nil
}

func (this *XsLeasingDealMadeType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsAddressTypeFuncValidate(&this.Address, ctx); err != nil { return err }
  if err := XsFactAddressTypeFuncValidate(&this.FactAddress, ctx); err != nil { return err }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsSoleProprietorTypeFuncValidate(&this.SoleProprietor, ctx); err != nil { return err }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncValidate(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncValidate(this.Application, ctx); err != nil { return err }
  }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncValidate(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncValidate(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncValidate(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncValidate(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncValidate(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsSourceNonMonetaryObligationTypeFuncValidate(&this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 1.5 Для принудительного исполнения передано требование о взыскании долга по алиментам, платы за жилое помещение, коммунальные услуги или услуги связи
type XsDebtClaimSubmittedTypeFunc func(this *XsDebtClaimSubmittedType, ctx *ParseContext) error

var XsDebtClaimSubmittedTypeFuncInit XsDebtClaimSubmittedTypeFunc = (*XsDebtClaimSubmittedType).Init

var XsDebtClaimSubmittedTypeFuncLoad XsDebtClaimSubmittedTypeFunc = (*XsDebtClaimSubmittedType).Load

var XsDebtClaimSubmittedTypeFuncValidate XsDebtClaimSubmittedTypeFunc = (*XsDebtClaimSubmittedType).Validate

func (this *XsDebtClaimSubmittedType) Init(ctx *ParseContext) error {
  if err := XsAddressTypeFuncInit(&this.Address, ctx); err != nil { return err }
  if err := XsFactAddressTypeFuncInit(&this.FactAddress, ctx); err != nil { return err }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsSoleProprietorTypeFuncInit(&this.SoleProprietor, ctx); err != nil { return err }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsCollectionTypeFuncInit(&this.Collection, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsDebtClaimSubmittedType) Load(ctx *ParseContext) error {
  if   !this.Address.Same(&ctx.person.Address) {
    ctx.person.Address = this.Address
  }
  if   !this.FactAddress.Same(ctx.person.FactAddress) {
    ctx.person.FactAddress = &this.FactAddress
  }
  foundContacts := make(map[int]bool)
  for i, v := range ctx.person.Contact {
    for k, w := range this.Contact {
      if w.Same(&v) {
        ctx.person.Contact[i] = w
        foundContacts[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Contact {
      if !foundContacts[i] {
        ctx.person.Contact = append(ctx.person.Contact, w)
      }
    }
  }
  if   !this.SoleProprietor.Same(ctx.person.SoleProprietor) {
    ctx.person.SoleProprietor = &this.SoleProprietor
  }
  foundLegalCapacitys := make(map[int]bool)
  for i, v := range ctx.person.LegalCapacity {
    for k, w := range this.LegalCapacity {
      if w.Same(&v) {
        ctx.person.LegalCapacity[i] = w
        foundLegalCapacitys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.LegalCapacity {
      if !foundLegalCapacitys[i] {
        ctx.person.LegalCapacity = append(ctx.person.LegalCapacity, w)
      }
    }
  }
  foundCollection := -1
  for i, v := range ctx.person.Collection {
    if this.Collection.Same(&v) {
      ctx.person.Collection[i] = this.Collection
      foundCollection = i
    }
  }
  if ctx.op != "D" && foundCollection == -1 {
    ctx.person.Collection = append(ctx.person.Collection, this.Collection)
  }
  if ctx.op == "D" && foundCollection != -1 {
    ctx.person.Collection = append(ctx.person.Collection[:foundCollection], ctx.person.Collection[foundCollection+1:]...)
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  return nil
}

func (this *XsDebtClaimSubmittedType) Validate(ctx *ParseContext) error {
  if err := XsAddressTypeFuncValidate(&this.Address, ctx); err != nil { return err }
  if err := XsFactAddressTypeFuncValidate(&this.FactAddress, ctx); err != nil { return err }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsSoleProprietorTypeFuncValidate(&this.SoleProprietor, ctx); err != nil { return err }
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsCollectionTypeFuncValidate(&this.Collection, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 1.6 Изменились сведения о взыскании долга по алиментам, платы за жилое помещение, коммунальные услуги или услуги связи либо указанное требование погашено
type XsDebtClaimChangedTypeFunc func(this *XsDebtClaimChangedType, ctx *ParseContext) error

var XsDebtClaimChangedTypeFuncInit XsDebtClaimChangedTypeFunc = (*XsDebtClaimChangedType).Init

var XsDebtClaimChangedTypeFuncLoad XsDebtClaimChangedTypeFunc = (*XsDebtClaimChangedType).Load

var XsDebtClaimChangedTypeFuncValidate XsDebtClaimChangedTypeFunc = (*XsDebtClaimChangedType).Validate

func (this *XsDebtClaimChangedType) Init(ctx *ParseContext) error {
  if err := XsCollectionTypeFuncInit(&this.Collection, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsDebtClaimChangedType) Load(ctx *ParseContext) error {
  foundCollection := -1
  for i, v := range ctx.person.Collection {
    if this.Collection.Same(&v) {
      ctx.person.Collection[i] = this.Collection
      foundCollection = i
    }
  }
  if ctx.op != "D" && foundCollection == -1 {
    ctx.person.Collection = append(ctx.person.Collection, this.Collection)
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  return nil
}

func (this *XsDebtClaimChangedType) Validate(ctx *ParseContext) error {
  if err := XsCollectionTypeFuncValidate(&this.Collection, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 1.7 и 1.8 Изменились сведения титульной части кредитной истории субъекта
type XsSubjectMainChangedTypeFunc func(this *XsSubjectMainChangedType, ctx *ParseContext) error

var XsSubjectMainChangedTypeFuncInit XsSubjectMainChangedTypeFunc = (*XsSubjectMainChangedType).Init

var XsSubjectMainChangedTypeFuncLoad XsSubjectMainChangedTypeFunc = (*XsSubjectMainChangedType).Load

var XsSubjectMainChangedTypeFuncValidate XsSubjectMainChangedTypeFunc = (*XsSubjectMainChangedType).Validate

func (this *XsSubjectMainChangedType) Init(ctx *ParseContext) error {
  if err := XsPersonNameTypeFuncInit(&this.Name, ctx); err != nil { return err }
  for _, v := range (*this).PrevName {
    if err := XsPersonPrevNameTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsBirthInfoTypeFuncInit(&this.BirthInfo, ctx); err != nil { return err }
  for _, v := range (*this).Id {
    if err := XsPersonIdTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevId {
    if err := XsPersonPrevIdTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Tax {
    if err := XsTaxTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.SocialId != nil {
    if err := XsSocialIdTypeFuncInit(this.SocialId, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsSubjectMainChangedType) Load(ctx *ParseContext) error {
  if   !this.Name.Same(&ctx.person.Name) {
    ctx.person.Name = this.Name
  }
  foundPrevNames := make(map[int]bool)
  for i, v := range ctx.person.PrevName {
    for k, w := range this.PrevName {
      if w.Same(&v) {
        ctx.person.PrevName[i] = w
        foundPrevNames[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevName {
      if !foundPrevNames[i] {
        ctx.person.PrevName = append(ctx.person.PrevName, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundPrevNames) > 0 {
      var newPrevNames []XsPersonPrevNameType
      for i, v := range ctx.person.PrevName {
        if !foundPrevNames[i] {
          newPrevNames = append(newPrevNames, v)
        }
      }
      ctx.person.PrevName = newPrevNames
    }
  }
  if   !this.BirthInfo.Same(&ctx.person.BirthInfo) {
    ctx.person.BirthInfo = this.BirthInfo
  }
  foundIds := make(map[int]bool)
  for i, v := range ctx.person.Id {
    for k, w := range this.Id {
      if w.Same(&v) {
        ctx.person.Id[i] = w
        foundIds[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Id {
      if !foundIds[i] {
        ctx.person.Id = append(ctx.person.Id, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundIds) > 0 {
      var newIds []XsPersonIdType
      for i, v := range ctx.person.Id {
        if !foundIds[i] {
          newIds = append(newIds, v)
        }
      }
      ctx.person.Id = newIds
    }
  }
  foundPrevIds := make(map[int]bool)
  for i, v := range ctx.person.PrevId {
    for k, w := range this.PrevId {
      if w.Same(&v) {
        ctx.person.PrevId[i] = w
        foundPrevIds[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevId {
      if !foundPrevIds[i] {
        ctx.person.PrevId = append(ctx.person.PrevId, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundPrevIds) > 0 {
      var newPrevIds []XsPersonPrevIdType
      for i, v := range ctx.person.PrevId {
        if !foundPrevIds[i] {
          newPrevIds = append(newPrevIds, v)
        }
      }
      ctx.person.PrevId = newPrevIds
    }
  }
  foundTaxs := make(map[int]bool)
  for i, v := range ctx.person.Tax {
    for k, w := range this.Tax {
      if w.Same(&v) {
        ctx.person.Tax[i] = w
        foundTaxs[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Tax {
      if !foundTaxs[i] {
        ctx.person.Tax = append(ctx.person.Tax, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundTaxs) > 0 {
      var newTaxs []XsTaxType
      for i, v := range ctx.person.Tax {
        if !foundTaxs[i] {
          newTaxs = append(newTaxs, v)
        }
      }
      ctx.person.Tax = newTaxs
    }
  }
  if  ctx.op == "D" && this.SocialId.Same(ctx.person.SocialId)  {
    ctx.person.SocialId = nil
  }
  if  ctx.op != "D" &&  !this.SocialId.Same(ctx.person.SocialId) {
    ctx.person.SocialId = this.SocialId
  }
  return nil
}

func (this *XsSubjectMainChangedType) Validate(ctx *ParseContext) error {
  if err := XsPersonNameTypeFuncValidate(&this.Name, ctx); err != nil { return err }
  for _, v := range (*this).PrevName {
    if err := XsPersonPrevNameTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsBirthInfoTypeFuncValidate(&this.BirthInfo, ctx); err != nil { return err }
  for _, v := range (*this).Id {
    if err := XsPersonIdTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevId {
    if err := XsPersonPrevIdTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Tax {
    if err := XsTaxTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.SocialId != nil {
    if err := XsSocialIdTypeFuncValidate(this.SocialId, ctx); err != nil { return err }
  }
  return nil
}

// 1.9 Изменились сведения о субъекте в основной части кредитной истории, кроме сведений о дееспособности, банкротстве, индивидуальном рейтинге и кредитной оценке
type XsSubjectSpecialChangedTypeFunc func(this *XsSubjectSpecialChangedType, ctx *ParseContext) error

var XsSubjectSpecialChangedTypeFuncInit XsSubjectSpecialChangedTypeFunc = (*XsSubjectSpecialChangedType).Init

var XsSubjectSpecialChangedTypeFuncLoad XsSubjectSpecialChangedTypeFunc = (*XsSubjectSpecialChangedType).Load

var XsSubjectSpecialChangedTypeFuncValidate XsSubjectSpecialChangedTypeFunc = (*XsSubjectSpecialChangedType).Validate

func (this *XsSubjectSpecialChangedType) Init(ctx *ParseContext) error {
  if err := XsAddressTypeFuncInit(&this.Address, ctx); err != nil { return err }
  if err := XsFactAddressTypeFuncInit(&this.FactAddress, ctx); err != nil { return err }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsSoleProprietorTypeFuncInit(&this.SoleProprietor, ctx); err != nil { return err }
  return nil
}

func (this *XsSubjectSpecialChangedType) Load(ctx *ParseContext) error {
  if   !this.Address.Same(&ctx.person.Address) {
    ctx.person.Address = this.Address
  }
  if  ctx.op == "D" && this.FactAddress.Same(ctx.person.FactAddress)  {
    ctx.person.FactAddress = nil
  }
  if  ctx.op != "D" &&  !this.FactAddress.Same(ctx.person.FactAddress) {
    ctx.person.FactAddress = &this.FactAddress
  }
  foundContacts := make(map[int]bool)
  for i, v := range ctx.person.Contact {
    for k, w := range this.Contact {
      if w.Same(&v) {
        ctx.person.Contact[i] = w
        foundContacts[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Contact {
      if !foundContacts[i] {
        ctx.person.Contact = append(ctx.person.Contact, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundContacts) > 0 {
      var newContacts []XsContactType
      for i, v := range ctx.person.Contact {
        if !foundContacts[i] {
          newContacts = append(newContacts, v)
        }
      }
      ctx.person.Contact = newContacts
    }
  }
  if  ctx.op == "D" && this.SoleProprietor.Same(ctx.person.SoleProprietor)  {
    ctx.person.SoleProprietor = nil
  }
  if  ctx.op != "D" &&  !this.SoleProprietor.Same(ctx.person.SoleProprietor) {
    ctx.person.SoleProprietor = &this.SoleProprietor
  }
  return nil
}

func (this *XsSubjectSpecialChangedType) Validate(ctx *ParseContext) error {
  if err := XsAddressTypeFuncValidate(&this.Address, ctx); err != nil { return err }
  if err := XsFactAddressTypeFuncValidate(&this.FactAddress, ctx); err != nil { return err }
  for _, v := range (*this).Contact {
    if err := XsContactTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsSoleProprietorTypeFuncValidate(&this.SoleProprietor, ctx); err != nil { return err }
  return nil
}

// 1.10 и 1.11 Изменились сведения о дееспособности субъекта
type XsSubjectCapacityChangedTypeFunc func(this *XsSubjectCapacityChangedType, ctx *ParseContext) error

var XsSubjectCapacityChangedTypeFuncInit XsSubjectCapacityChangedTypeFunc = (*XsSubjectCapacityChangedType).Init

var XsSubjectCapacityChangedTypeFuncLoad XsSubjectCapacityChangedTypeFunc = (*XsSubjectCapacityChangedType).Load

var XsSubjectCapacityChangedTypeFuncValidate XsSubjectCapacityChangedTypeFunc = (*XsSubjectCapacityChangedType).Validate

func (this *XsSubjectCapacityChangedType) Init(ctx *ParseContext) error {
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsSubjectCapacityChangedType) Load(ctx *ParseContext) error {
  foundLegalCapacitys := make(map[int]bool)
  for i, v := range ctx.person.LegalCapacity {
    for k, w := range this.LegalCapacity {
      if w.Same(&v) {
        ctx.person.LegalCapacity[i] = w
        foundLegalCapacitys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.LegalCapacity {
      if !foundLegalCapacitys[i] {
        ctx.person.LegalCapacity = append(ctx.person.LegalCapacity, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundLegalCapacitys) > 0 {
      var newLegalCapacitys []XsLegalCapacityType
      for i, v := range ctx.person.LegalCapacity {
        if !foundLegalCapacitys[i] {
          newLegalCapacitys = append(newLegalCapacitys, v)
        }
      }
      ctx.person.LegalCapacity = newLegalCapacitys
    }
  }
  return nil
}

func (this *XsSubjectCapacityChangedType) Validate(ctx *ParseContext) error {
  for _, v := range (*this).LegalCapacity {
    if err := XsLegalCapacityTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// 1.12 Изменились сведения по делу о банкротстве субъекта
type XsSubjectBankruptcyChangedTypeFunc func(this *XsSubjectBankruptcyChangedType, ctx *ParseContext) error

var XsSubjectBankruptcyChangedTypeFuncInit XsSubjectBankruptcyChangedTypeFunc = (*XsSubjectBankruptcyChangedType).Init

var XsSubjectBankruptcyChangedTypeFuncLoad XsSubjectBankruptcyChangedTypeFunc = (*XsSubjectBankruptcyChangedType).Load

var XsSubjectBankruptcyChangedTypeFuncValidate XsSubjectBankruptcyChangedTypeFunc = (*XsSubjectBankruptcyChangedType).Validate

func (this *XsSubjectBankruptcyChangedType) Init(ctx *ParseContext) error {
  if err := XsBankruptcyTypeFuncInit(&this.Bankruptcy, ctx); err != nil { return err }
  if err := XsFulfillmentTypeFuncInit(&this.Fulfillment, ctx); err != nil { return err }
  return nil
}

func (this *XsSubjectBankruptcyChangedType) Load(ctx *ParseContext) error {
  foundBankruptcy := -1
  for i, v := range ctx.person.Bankruptcy {
    if this.Bankruptcy.Same(&v) {
      ctx.person.Bankruptcy[i] = this.Bankruptcy
      foundBankruptcy = i
    }
  }
  if ctx.op != "D" && foundBankruptcy == -1 {
    ctx.person.Bankruptcy = append(ctx.person.Bankruptcy, this.Bankruptcy)
  }
  foundFulfillment := -1
  for i, v := range ctx.person.Fulfillment {
    if this.Fulfillment.Same(&v) {
      ctx.person.Fulfillment[i] = this.Fulfillment
      foundFulfillment = i
    }
  }
  if ctx.op != "D" && foundFulfillment == -1 {
    ctx.person.Fulfillment = append(ctx.person.Fulfillment, this.Fulfillment)
  }
  if ctx.op == "D" && foundFulfillment != -1 {
    ctx.person.Fulfillment = append(ctx.person.Fulfillment[:foundFulfillment], ctx.person.Fulfillment[foundFulfillment+1:]...)
  }
  return nil
}

func (this *XsSubjectBankruptcyChangedType) Validate(ctx *ParseContext) error {
  if err := XsBankruptcyTypeFuncValidate(&this.Bankruptcy, ctx); err != nil { return err }
  if err := XsFulfillmentTypeFuncValidate(&this.Fulfillment, ctx); err != nil { return err }
  return nil
}

// 1.13 Рассчитан индивидуальный рейтинг субъекта (вследствие обращения за его рейтингом или кредитным отчетом)
type XsRatingCalculatedTypeFunc func(this *XsRatingCalculatedType, ctx *ParseContext) error

var XsRatingCalculatedTypeFuncInit XsRatingCalculatedTypeFunc = (*XsRatingCalculatedType).Init

var XsRatingCalculatedTypeFuncLoad XsRatingCalculatedTypeFunc = (*XsRatingCalculatedType).Load

var XsRatingCalculatedTypeFuncValidate XsRatingCalculatedTypeFunc = (*XsRatingCalculatedType).Validate

func (this *XsRatingCalculatedType) Init(ctx *ParseContext) error {
  if err := XsRatingTypeFuncInit(&this.Rating, ctx); err != nil { return err }
  return nil
}

func (this *XsRatingCalculatedType) Load(ctx *ParseContext) error {
  if err := XsRatingTypeFuncLoad(&this.Rating, ctx); err != nil { return err }
  return nil
}

func (this *XsRatingCalculatedType) Validate(ctx *ParseContext) error {
  if err := XsRatingTypeFuncValidate(&this.Rating, ctx); err != nil { return err }
  return nil
}

// 1.14 Пользователь обратился за кредитной оценкой (скорингом) субъекта
type XsScoringRequestedTypeFunc func(this *XsScoringRequestedType, ctx *ParseContext) error

var XsScoringRequestedTypeFuncInit XsScoringRequestedTypeFunc = (*XsScoringRequestedType).Init

var XsScoringRequestedTypeFuncLoad XsScoringRequestedTypeFunc = (*XsScoringRequestedType).Load

var XsScoringRequestedTypeFuncValidate XsScoringRequestedTypeFunc = (*XsScoringRequestedType).Validate

func (this *XsScoringRequestedType) Init(ctx *ParseContext) error {
  if err := XsScoreTypeFuncInit(&this.Score, ctx); err != nil { return err }
  return nil
}

func (this *XsScoringRequestedType) Load(ctx *ParseContext) error {
  if err := XsScoreTypeFuncLoad(&this.Score, ctx); err != nil { return err }
  return nil
}

func (this *XsScoringRequestedType) Validate(ctx *ParseContext) error {
  if err := XsScoreTypeFuncValidate(&this.Score, ctx); err != nil { return err }
  return nil
}

// 1.15 Пользователь запросил кредитный отчет субъекта
type XsReportRequestedTypeFunc func(this *XsReportRequestedType, ctx *ParseContext) error

var XsReportRequestedTypeFuncInit XsReportRequestedTypeFunc = (*XsReportRequestedType).Init

var XsReportRequestedTypeFuncLoad XsReportRequestedTypeFunc = (*XsReportRequestedType).Load

var XsReportRequestedTypeFuncValidate XsReportRequestedTypeFunc = (*XsReportRequestedType).Validate

func (this *XsReportRequestedType) Init(ctx *ParseContext) error {
  if err := XsCustomerRequestTypeFuncInit(&this.CustomerRequest, ctx); err != nil { return err }
  return nil
}

func (this *XsReportRequestedType) Load(ctx *ParseContext) error {
  if err := XsCustomerRequestTypeFuncLoad(&this.CustomerRequest, ctx); err != nil { return err }
  return nil
}

func (this *XsReportRequestedType) Validate(ctx *ParseContext) error {
  if err := XsCustomerRequestTypeFuncValidate(&this.CustomerRequest, ctx); err != nil { return err }
  return nil
}

// 2.1 Изменились сведения об условиях обязательства субъекта для денежного обязательства
type XsMonetaryDealChangedTypeFunc func(this *XsMonetaryDealChangedType, ctx *ParseContext) error

var XsMonetaryDealChangedTypeFuncInit XsMonetaryDealChangedTypeFunc = (*XsMonetaryDealChangedType).Init

var XsMonetaryDealChangedTypeFuncLoad XsMonetaryDealChangedTypeFunc = (*XsMonetaryDealChangedType).Load

var XsMonetaryDealChangedTypeFuncValidate XsMonetaryDealChangedTypeFunc = (*XsMonetaryDealChangedType).Validate

func (this *XsMonetaryDealChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncInit(&this.JointDebtors, ctx); err != nil { return err }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncInit(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncInit(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncInit(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncInit(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncInit(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncInit(this.TotalCost, ctx); err != nil { return err }
  }
  if this.MonthlyPayment != nil {
    if err := XsMonthlyPaymentTypeFuncInit(this.MonthlyPayment, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsMonetaryDealChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if   !this.JointDebtors.Same(ctx.personDeal.JointDebtors) {
    ctx.personDeal.JointDebtors = &this.JointDebtors
  }
  foundChanges := make(map[int]bool)
  for i, v := range ctx.personDeal.Change {
    for k, w := range this.Change {
      if w.Same(&v) {
        ctx.personDeal.Change[i] = w
        foundChanges[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Change {
      if !foundChanges[i] {
        ctx.personDeal.Change = append(ctx.personDeal.Change, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundChanges) > 0 {
      var newChanges []XsDealChangeType
      for i, v := range ctx.personDeal.Change {
        if !foundChanges[i] {
          newChanges = append(newChanges, v)
        }
      }
      ctx.personDeal.Change = newChanges
    }
  }
  if   !this.Accounting.Same(ctx.personDeal.Accounting) {
    ctx.personDeal.Accounting = &this.Accounting
  }
  if   !this.Amount.Same(ctx.personDeal.Amount) {
    ctx.personDeal.Amount = &this.Amount
  }
  if   !this.PaymentTerms.Same(ctx.personDeal.PaymentTerms) {
    ctx.personDeal.PaymentTerms = &this.PaymentTerms
  }
  if   !this.FundDate.Same(ctx.personDeal.FundDate) {
    ctx.personDeal.FundDate = &this.FundDate
  }
  foundArrear := -1
  for i, v := range ctx.personDeal.Arrear {
    if this.Arrear.Same(&v) {
      ctx.personDeal.Arrear[i] = this.Arrear
      foundArrear = i
    }
  }
  if ctx.op != "D" && foundArrear == -1 {
    ctx.personDeal.Arrear = append(ctx.personDeal.Arrear, this.Arrear)
  }
  if ctx.op == "D" && foundArrear != -1 {
    ctx.personDeal.Arrear = append(ctx.personDeal.Arrear[:foundArrear], ctx.personDeal.Arrear[foundArrear+1:]...)
  }
  foundDueArrear := -1
  for i, v := range ctx.personDeal.DueArrear {
    if this.DueArrear.Same(&v) {
      ctx.personDeal.DueArrear[i] = this.DueArrear
      foundDueArrear = i
    }
  }
  if ctx.op != "D" && foundDueArrear == -1 {
    ctx.personDeal.DueArrear = append(ctx.personDeal.DueArrear, this.DueArrear)
  }
  if ctx.op == "D" && foundDueArrear != -1 {
    ctx.personDeal.DueArrear = append(ctx.personDeal.DueArrear[:foundDueArrear], ctx.personDeal.DueArrear[foundDueArrear+1:]...)
  }
  foundPastdueArrear := -1
  for i, v := range ctx.personDeal.PastdueArrear {
    if this.PastdueArrear.Same(&v) {
      ctx.personDeal.PastdueArrear[i] = this.PastdueArrear
      foundPastdueArrear = i
    }
  }
  if ctx.op != "D" && foundPastdueArrear == -1 {
    ctx.personDeal.PastdueArrear = append(ctx.personDeal.PastdueArrear, this.PastdueArrear)
  }
  foundPaymentss := make(map[int]bool)
  for i, v := range ctx.personDeal.Payments {
    for k, w := range this.Payments {
      if w.Same(&v) {
        ctx.personDeal.Payments[i] = w
        foundPaymentss[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Payments {
      if !foundPaymentss[i] {
        ctx.personDeal.Payments = append(ctx.personDeal.Payments, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundPaymentss) > 0 {
      var newPaymentss []XsPaymentsType
      for i, v := range ctx.personDeal.Payments {
        if !foundPaymentss[i] {
          newPaymentss = append(newPaymentss, v)
        }
      }
      ctx.personDeal.Payments = newPaymentss
    }
  }
  if  ctx.op == "D" && this.TotalCost.Same(ctx.personDeal.TotalCost)  {
    ctx.personDeal.TotalCost = nil
  }
  if  ctx.op != "D" &&  !this.TotalCost.Same(ctx.personDeal.TotalCost) {
    ctx.personDeal.TotalCost = this.TotalCost
  }
  if  ctx.op == "D" && this.MonthlyPayment.Same(ctx.personDeal.MonthlyPayment)  {
    ctx.personDeal.MonthlyPayment = nil
  }
  if  ctx.op != "D" &&  !this.MonthlyPayment.Same(ctx.personDeal.MonthlyPayment) {
    ctx.personDeal.MonthlyPayment = this.MonthlyPayment
  }
  return nil
}

func (this *XsMonetaryDealChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncValidate(&this.JointDebtors, ctx); err != nil { return err }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncValidate(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncValidate(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncValidate(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncValidate(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncValidate(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncValidate(this.TotalCost, ctx); err != nil { return err }
  }
  if this.MonthlyPayment != nil {
    if err := XsMonthlyPaymentTypeFuncValidate(this.MonthlyPayment, ctx); err != nil { return err }
  }
  return nil
}

// 2.1 Изменились сведения об условиях обязательства субъекта для неденежного обязательства
type XsNonMonetaryDealChangedTypeFunc func(this *XsNonMonetaryDealChangedType, ctx *ParseContext) error

var XsNonMonetaryDealChangedTypeFuncInit XsNonMonetaryDealChangedTypeFunc = (*XsNonMonetaryDealChangedType).Init

var XsNonMonetaryDealChangedTypeFuncLoad XsNonMonetaryDealChangedTypeFunc = (*XsNonMonetaryDealChangedType).Load

var XsNonMonetaryDealChangedTypeFuncValidate XsNonMonetaryDealChangedTypeFunc = (*XsNonMonetaryDealChangedType).Validate

func (this *XsNonMonetaryDealChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncInit(&this.JointDebtors, ctx); err != nil { return err }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if err := XsSubjectNonMonetaryObligationTypeFuncInit(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

func (this *XsNonMonetaryDealChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if   !this.JointDebtors.Same(ctx.personDeal.JointDebtors) {
    ctx.personDeal.JointDebtors = &this.JointDebtors
  }
  foundChanges := make(map[int]bool)
  for i, v := range ctx.personDeal.Change {
    for k, w := range this.Change {
      if w.Same(&v) {
        ctx.personDeal.Change[i] = w
        foundChanges[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Change {
      if !foundChanges[i] {
        ctx.personDeal.Change = append(ctx.personDeal.Change, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundChanges) > 0 {
      var newChanges []XsDealChangeType
      for i, v := range ctx.personDeal.Change {
        if !foundChanges[i] {
          newChanges = append(newChanges, v)
        }
      }
      ctx.personDeal.Change = newChanges
    }
  }
  if   !this.Accounting.Same(ctx.personDeal.Accounting) {
    ctx.personDeal.Accounting = &this.Accounting
  }
  if   !this.SubjectNonMonetaryObligation.Same(ctx.personDeal.SubjectNonMonetaryObligation) {
    ctx.personDeal.SubjectNonMonetaryObligation = &this.SubjectNonMonetaryObligation
  }
  return nil
}

func (this *XsNonMonetaryDealChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncValidate(&this.JointDebtors, ctx); err != nil { return err }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if err := XsSubjectNonMonetaryObligationTypeFuncValidate(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

// 2.2 Субъекту передана сумма займа (кредита) либо субъект стал принципалом по гарантии или поручителем по сделке, кроме договора лизинга для денежного обязательства
type XsMonetaryDealPrincipalChangedTypeFunc func(this *XsMonetaryDealPrincipalChangedType, ctx *ParseContext) error

var XsMonetaryDealPrincipalChangedTypeFuncInit XsMonetaryDealPrincipalChangedTypeFunc = (*XsMonetaryDealPrincipalChangedType).Init

var XsMonetaryDealPrincipalChangedTypeFuncLoad XsMonetaryDealPrincipalChangedTypeFunc = (*XsMonetaryDealPrincipalChangedType).Load

var XsMonetaryDealPrincipalChangedTypeFuncValidate XsMonetaryDealPrincipalChangedTypeFunc = (*XsMonetaryDealPrincipalChangedType).Validate

func (this *XsMonetaryDealPrincipalChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncInit(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncInit(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncInit(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncInit(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncInit(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncInit(this.TotalCost, ctx); err != nil { return err }
  }
  if this.MonthlyPayment != nil {
    if err := XsMonthlyPaymentTypeFuncInit(this.MonthlyPayment, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsMonetaryDealPrincipalChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  if   !this.Amount.Same(ctx.personDeal.Amount) {
    ctx.personDeal.Amount = &this.Amount
  }
  if   !this.PaymentTerms.Same(ctx.personDeal.PaymentTerms) {
    ctx.personDeal.PaymentTerms = &this.PaymentTerms
  }
  if   !this.FundDate.Same(ctx.personDeal.FundDate) {
    ctx.personDeal.FundDate = &this.FundDate
  }
  foundArrear := -1
  for i, v := range ctx.personDeal.Arrear {
    if this.Arrear.Same(&v) {
      ctx.personDeal.Arrear[i] = this.Arrear
      foundArrear = i
    }
  }
  if ctx.op != "D" && foundArrear == -1 {
    ctx.personDeal.Arrear = append(ctx.personDeal.Arrear, this.Arrear)
  }
  foundDueArrear := -1
  for i, v := range ctx.personDeal.DueArrear {
    if this.DueArrear.Same(&v) {
      ctx.personDeal.DueArrear[i] = this.DueArrear
      foundDueArrear = i
    }
  }
  if ctx.op != "D" && foundDueArrear == -1 {
    ctx.personDeal.DueArrear = append(ctx.personDeal.DueArrear, this.DueArrear)
  }
  foundPastdueArrear := -1
  for i, v := range ctx.personDeal.PastdueArrear {
    if this.PastdueArrear.Same(&v) {
      ctx.personDeal.PastdueArrear[i] = this.PastdueArrear
      foundPastdueArrear = i
    }
  }
  if ctx.op != "D" && foundPastdueArrear == -1 {
    ctx.personDeal.PastdueArrear = append(ctx.personDeal.PastdueArrear, this.PastdueArrear)
  }
  foundPaymentss := make(map[int]bool)
  for i, v := range ctx.personDeal.Payments {
    for k, w := range this.Payments {
      if w.Same(&v) {
        ctx.personDeal.Payments[i] = w
        foundPaymentss[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Payments {
      if !foundPaymentss[i] {
        ctx.personDeal.Payments = append(ctx.personDeal.Payments, w)
      }
    }
  }
  if   !this.TotalCost.Same(ctx.personDeal.TotalCost) {
    ctx.personDeal.TotalCost = this.TotalCost
  }
  if   !this.MonthlyPayment.Same(ctx.personDeal.MonthlyPayment) {
    ctx.personDeal.MonthlyPayment = this.MonthlyPayment
  }
  return nil
}

func (this *XsMonetaryDealPrincipalChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncValidate(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncValidate(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncValidate(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncValidate(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncValidate(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncValidate(this.TotalCost, ctx); err != nil { return err }
  }
  if this.MonthlyPayment != nil {
    if err := XsMonthlyPaymentTypeFuncValidate(this.MonthlyPayment, ctx); err != nil { return err }
  }
  return nil
}

// 2.2 Субъекту передана сумма займа (кредита) либо субъект стал принципалом по гарантии или поручителем по сделке, кроме договора лизинга для неденежного обязательства
type XsNonMonetaryDealPrincipalChangedTypeFunc func(this *XsNonMonetaryDealPrincipalChangedType, ctx *ParseContext) error

var XsNonMonetaryDealPrincipalChangedTypeFuncInit XsNonMonetaryDealPrincipalChangedTypeFunc = (*XsNonMonetaryDealPrincipalChangedType).Init

var XsNonMonetaryDealPrincipalChangedTypeFuncLoad XsNonMonetaryDealPrincipalChangedTypeFunc = (*XsNonMonetaryDealPrincipalChangedType).Load

var XsNonMonetaryDealPrincipalChangedTypeFuncValidate XsNonMonetaryDealPrincipalChangedTypeFunc = (*XsNonMonetaryDealPrincipalChangedType).Validate

func (this *XsNonMonetaryDealPrincipalChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  if err := XsSourceNonMonetaryObligationTypeFuncInit(&this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

func (this *XsNonMonetaryDealPrincipalChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  if   !this.SourceNonMonetaryObligation.Same(ctx.personDeal.SourceNonMonetaryObligation) {
    ctx.personDeal.SourceNonMonetaryObligation = &this.SourceNonMonetaryObligation
  }
  return nil
}

func (this *XsNonMonetaryDealPrincipalChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  if err := XsSourceNonMonetaryObligationTypeFuncValidate(&this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

// 2.3 Изменились сведения об исполнении обязательства субъектом, наступила ответственность поручителя или обязательство принципала возместить выплаченную сумму для денежного обязательства
type XsMonetaryDealPerformanceChangedTypeFunc func(this *XsMonetaryDealPerformanceChangedType, ctx *ParseContext) error

var XsMonetaryDealPerformanceChangedTypeFuncInit XsMonetaryDealPerformanceChangedTypeFunc = (*XsMonetaryDealPerformanceChangedType).Init

var XsMonetaryDealPerformanceChangedTypeFuncLoad XsMonetaryDealPerformanceChangedTypeFunc = (*XsMonetaryDealPerformanceChangedType).Load

var XsMonetaryDealPerformanceChangedTypeFuncValidate XsMonetaryDealPerformanceChangedTypeFunc = (*XsMonetaryDealPerformanceChangedType).Validate

func (this *XsMonetaryDealPerformanceChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncInit(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncInit(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncInit(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncInit(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncInit(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncInit(this.TotalCost, ctx); err != nil { return err }
  }
  if this.MonthlyPayment != nil {
    if err := XsMonthlyPaymentTypeFuncInit(this.MonthlyPayment, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsMonetaryDealPerformanceChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if   !this.Accounting.Same(ctx.personDeal.Accounting) {
    ctx.personDeal.Accounting = &this.Accounting
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  if ctx.op == "D" && foundParticipation != -1 {
    ctx.person.Participation = append(ctx.person.Participation[:foundParticipation], ctx.person.Participation[foundParticipation+1:]...)
  }
  if   !this.Amount.Same(ctx.personDeal.Amount) {
    ctx.personDeal.Amount = &this.Amount
  }
  if   !this.PaymentTerms.Same(ctx.personDeal.PaymentTerms) {
    ctx.personDeal.PaymentTerms = &this.PaymentTerms
  }
  if   !this.FundDate.Same(ctx.personDeal.FundDate) {
    ctx.personDeal.FundDate = &this.FundDate
  }
  foundArrear := -1
  for i, v := range ctx.personDeal.Arrear {
    if this.Arrear.Same(&v) {
      ctx.personDeal.Arrear[i] = this.Arrear
      foundArrear = i
    }
  }
  if ctx.op != "D" && foundArrear == -1 {
    ctx.personDeal.Arrear = append(ctx.personDeal.Arrear, this.Arrear)
  }
  foundDueArrear := -1
  for i, v := range ctx.personDeal.DueArrear {
    if this.DueArrear.Same(&v) {
      ctx.personDeal.DueArrear[i] = this.DueArrear
      foundDueArrear = i
    }
  }
  if ctx.op != "D" && foundDueArrear == -1 {
    ctx.personDeal.DueArrear = append(ctx.personDeal.DueArrear, this.DueArrear)
  }
  foundPastdueArrear := -1
  for i, v := range ctx.personDeal.PastdueArrear {
    if this.PastdueArrear.Same(&v) {
      ctx.personDeal.PastdueArrear[i] = this.PastdueArrear
      foundPastdueArrear = i
    }
  }
  if ctx.op != "D" && foundPastdueArrear == -1 {
    ctx.personDeal.PastdueArrear = append(ctx.personDeal.PastdueArrear, this.PastdueArrear)
  }
  foundPaymentss := make(map[int]bool)
  for i, v := range ctx.personDeal.Payments {
    for k, w := range this.Payments {
      if w.Same(&v) {
        ctx.personDeal.Payments[i] = w
        foundPaymentss[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Payments {
      if !foundPaymentss[i] {
        ctx.personDeal.Payments = append(ctx.personDeal.Payments, w)
      }
    }
  }
  if   !this.TotalCost.Same(ctx.personDeal.TotalCost) {
    ctx.personDeal.TotalCost = this.TotalCost
  }
  if   !this.MonthlyPayment.Same(ctx.personDeal.MonthlyPayment) {
    ctx.personDeal.MonthlyPayment = this.MonthlyPayment
  }
  return nil
}

func (this *XsMonetaryDealPerformanceChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncValidate(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncValidate(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncValidate(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncValidate(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncValidate(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncValidate(this.TotalCost, ctx); err != nil { return err }
  }
  if this.MonthlyPayment != nil {
    if err := XsMonthlyPaymentTypeFuncValidate(this.MonthlyPayment, ctx); err != nil { return err }
  }
  return nil
}

// 2.3 Изменились сведения об исполнении обязательства субъектом, наступила ответственность поручителя или обязательство принципала возместить выплаченную сумму для неденежного обязательства
type XsNonMonetaryDealPerformanceChangedTypeFunc func(this *XsNonMonetaryDealPerformanceChangedType, ctx *ParseContext) error

var XsNonMonetaryDealPerformanceChangedTypeFuncInit XsNonMonetaryDealPerformanceChangedTypeFunc = (*XsNonMonetaryDealPerformanceChangedType).Init

var XsNonMonetaryDealPerformanceChangedTypeFuncLoad XsNonMonetaryDealPerformanceChangedTypeFunc = (*XsNonMonetaryDealPerformanceChangedType).Load

var XsNonMonetaryDealPerformanceChangedTypeFuncValidate XsNonMonetaryDealPerformanceChangedTypeFunc = (*XsNonMonetaryDealPerformanceChangedType).Validate

func (this *XsNonMonetaryDealPerformanceChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  if err := XsSubjectNonMonetaryObligationTypeFuncInit(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

func (this *XsNonMonetaryDealPerformanceChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if   !this.Accounting.Same(ctx.personDeal.Accounting) {
    ctx.personDeal.Accounting = &this.Accounting
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  if ctx.op == "D" && foundParticipation != -1 {
    ctx.person.Participation = append(ctx.person.Participation[:foundParticipation], ctx.person.Participation[foundParticipation+1:]...)
  }
  if   !this.SubjectNonMonetaryObligation.Same(ctx.personDeal.SubjectNonMonetaryObligation) {
    ctx.personDeal.SubjectNonMonetaryObligation = &this.SubjectNonMonetaryObligation
  }
  return nil
}

func (this *XsNonMonetaryDealPerformanceChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  if err := XsSubjectNonMonetaryObligationTypeFuncValidate(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

// 2.4 Изменились сведения об обеспечении исполнения обязательства
type XsDealSecuringChangedTypeFunc func(this *XsDealSecuringChangedType, ctx *ParseContext) error

var XsDealSecuringChangedTypeFuncInit XsDealSecuringChangedTypeFunc = (*XsDealSecuringChangedType).Init

var XsDealSecuringChangedTypeFuncLoad XsDealSecuringChangedTypeFunc = (*XsDealSecuringChangedType).Load

var XsDealSecuringChangedTypeFuncValidate XsDealSecuringChangedTypeFunc = (*XsDealSecuringChangedType).Validate

func (this *XsDealSecuringChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
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
  return nil
}

func (this *XsDealSecuringChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  foundCollaterals := make(map[int]bool)
  for i, v := range ctx.personDeal.Collateral {
    for k, w := range this.Collateral {
      if w.Same(&v) {
        ctx.personDeal.Collateral[i] = w
        foundCollaterals[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Collateral {
      if !foundCollaterals[i] {
        ctx.personDeal.Collateral = append(ctx.personDeal.Collateral, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundCollaterals) > 0 {
      var newCollaterals []XsCollateralType
      for i, v := range ctx.personDeal.Collateral {
        if !foundCollaterals[i] {
          newCollaterals = append(newCollaterals, v)
        }
      }
      ctx.personDeal.Collateral = newCollaterals
    }
  }
  foundWarrantys := make(map[int]bool)
  for i, v := range ctx.personDeal.Warranty {
    for k, w := range this.Warranty {
      if w.Same(&v) {
        ctx.personDeal.Warranty[i] = w
        foundWarrantys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Warranty {
      if !foundWarrantys[i] {
        ctx.personDeal.Warranty = append(ctx.personDeal.Warranty, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundWarrantys) > 0 {
      var newWarrantys []XsWarrantyType
      for i, v := range ctx.personDeal.Warranty {
        if !foundWarrantys[i] {
          newWarrantys = append(newWarrantys, v)
        }
      }
      ctx.personDeal.Warranty = newWarrantys
    }
  }
  foundIndepGuarantees := make(map[int]bool)
  for i, v := range ctx.personDeal.IndepGuarantee {
    for k, w := range this.IndepGuarantee {
      if w.Same(&v) {
        ctx.personDeal.IndepGuarantee[i] = w
        foundIndepGuarantees[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.IndepGuarantee {
      if !foundIndepGuarantees[i] {
        ctx.personDeal.IndepGuarantee = append(ctx.personDeal.IndepGuarantee, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundIndepGuarantees) > 0 {
      var newIndepGuarantees []XsIndepGuaranteeType
      for i, v := range ctx.personDeal.IndepGuarantee {
        if !foundIndepGuarantees[i] {
          newIndepGuarantees = append(newIndepGuarantees, v)
        }
      }
      ctx.personDeal.IndepGuarantee = newIndepGuarantees
    }
  }
  foundCollateralInsurances := make(map[int]bool)
  for i, v := range ctx.personDeal.CollateralInsurance {
    for k, w := range this.CollateralInsurance {
      if w.Same(&v) {
        ctx.personDeal.CollateralInsurance[i] = w
        foundCollateralInsurances[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.CollateralInsurance {
      if !foundCollateralInsurances[i] {
        ctx.personDeal.CollateralInsurance = append(ctx.personDeal.CollateralInsurance, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundCollateralInsurances) > 0 {
      var newCollateralInsurances []XsCollateralInsuranceType
      for i, v := range ctx.personDeal.CollateralInsurance {
        if !foundCollateralInsurances[i] {
          newCollateralInsurances = append(newCollateralInsurances, v)
        }
      }
      ctx.personDeal.CollateralInsurance = newCollateralInsurances
    }
  }
  foundSettlements := make(map[int]bool)
  for i, v := range ctx.personDeal.Settlement {
    for k, w := range this.Settlement {
      if w.Same(&v) {
        ctx.personDeal.Settlement[i] = w
        foundSettlements[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Settlement {
      if !foundSettlements[i] {
        ctx.personDeal.Settlement = append(ctx.personDeal.Settlement, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundSettlements) > 0 {
      var newSettlements []XsSettlementType
      for i, v := range ctx.personDeal.Settlement {
        if !foundSettlements[i] {
          newSettlements = append(newSettlements, v)
        }
      }
      ctx.personDeal.Settlement = newSettlements
    }
  }
  if  ctx.op == "D" && this.Repayment.Same(ctx.personDeal.Repayment)  {
    ctx.personDeal.Repayment = nil
  }
  if  ctx.op != "D" &&  !this.Repayment.Same(ctx.personDeal.Repayment) {
    ctx.personDeal.Repayment = this.Repayment
  }
  return nil
}

func (this *XsDealSecuringChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
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
  return nil
}

// 2.5 Обязательство субъекта прекратилось для денежного обязательства
type XsMonetaryDealEndedTypeFunc func(this *XsMonetaryDealEndedType, ctx *ParseContext) error

var XsMonetaryDealEndedTypeFuncInit XsMonetaryDealEndedTypeFunc = (*XsMonetaryDealEndedType).Init

var XsMonetaryDealEndedTypeFuncLoad XsMonetaryDealEndedTypeFunc = (*XsMonetaryDealEndedType).Load

var XsMonetaryDealEndedTypeFuncValidate XsMonetaryDealEndedTypeFunc = (*XsMonetaryDealEndedType).Validate

func (this *XsMonetaryDealEndedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsTerminationTypeFuncInit(&this.Termination, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncInit(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncInit(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncInit(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncInit(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncInit(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncInit(this.TotalCost, ctx); err != nil { return err }
  }
  if this.MonthlyPayment != nil {
    if err := XsMonthlyPaymentTypeFuncInit(this.MonthlyPayment, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsMonetaryDealEndedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if  ctx.op == "D" && this.Termination.Same(ctx.personDeal.Termination)  {
    ctx.personDeal.Termination = nil
  }
  if  ctx.op != "D" &&  !this.Termination.Same(ctx.personDeal.Termination) {
    ctx.personDeal.Termination = &this.Termination
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  if   !this.Amount.Same(ctx.personDeal.Amount) {
    ctx.personDeal.Amount = &this.Amount
  }
  if   !this.PaymentTerms.Same(ctx.personDeal.PaymentTerms) {
    ctx.personDeal.PaymentTerms = &this.PaymentTerms
  }
  if   !this.FundDate.Same(ctx.personDeal.FundDate) {
    ctx.personDeal.FundDate = &this.FundDate
  }
  foundArrear := -1
  for i, v := range ctx.personDeal.Arrear {
    if this.Arrear.Same(&v) {
      ctx.personDeal.Arrear[i] = this.Arrear
      foundArrear = i
    }
  }
  if ctx.op != "D" && foundArrear == -1 {
    ctx.personDeal.Arrear = append(ctx.personDeal.Arrear, this.Arrear)
  }
  foundDueArrear := -1
  for i, v := range ctx.personDeal.DueArrear {
    if this.DueArrear.Same(&v) {
      ctx.personDeal.DueArrear[i] = this.DueArrear
      foundDueArrear = i
    }
  }
  if ctx.op != "D" && foundDueArrear == -1 {
    ctx.personDeal.DueArrear = append(ctx.personDeal.DueArrear, this.DueArrear)
  }
  foundPastdueArrear := -1
  for i, v := range ctx.personDeal.PastdueArrear {
    if this.PastdueArrear.Same(&v) {
      ctx.personDeal.PastdueArrear[i] = this.PastdueArrear
      foundPastdueArrear = i
    }
  }
  if ctx.op != "D" && foundPastdueArrear == -1 {
    ctx.personDeal.PastdueArrear = append(ctx.personDeal.PastdueArrear, this.PastdueArrear)
  }
  foundPaymentss := make(map[int]bool)
  for i, v := range ctx.personDeal.Payments {
    for k, w := range this.Payments {
      if w.Same(&v) {
        ctx.personDeal.Payments[i] = w
        foundPaymentss[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Payments {
      if !foundPaymentss[i] {
        ctx.personDeal.Payments = append(ctx.personDeal.Payments, w)
      }
    }
  }
  if   !this.TotalCost.Same(ctx.personDeal.TotalCost) {
    ctx.personDeal.TotalCost = this.TotalCost
  }
  if   !this.MonthlyPayment.Same(ctx.personDeal.MonthlyPayment) {
    ctx.personDeal.MonthlyPayment = this.MonthlyPayment
  }
  return nil
}

func (this *XsMonetaryDealEndedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsTerminationTypeFuncValidate(&this.Termination, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncValidate(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncValidate(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncValidate(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncValidate(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncValidate(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.TotalCost != nil {
    if err := XsTotalCostTypeFuncValidate(this.TotalCost, ctx); err != nil { return err }
  }
  if this.MonthlyPayment != nil {
    if err := XsMonthlyPaymentTypeFuncValidate(this.MonthlyPayment, ctx); err != nil { return err }
  }
  return nil
}

// 2.5 Обязательство субъекта прекратилось для неденежного обязательства
type XsNonMonetaryDealEndedTypeFunc func(this *XsNonMonetaryDealEndedType, ctx *ParseContext) error

var XsNonMonetaryDealEndedTypeFuncInit XsNonMonetaryDealEndedTypeFunc = (*XsNonMonetaryDealEndedType).Init

var XsNonMonetaryDealEndedTypeFuncLoad XsNonMonetaryDealEndedTypeFunc = (*XsNonMonetaryDealEndedType).Load

var XsNonMonetaryDealEndedTypeFuncValidate XsNonMonetaryDealEndedTypeFunc = (*XsNonMonetaryDealEndedType).Validate

func (this *XsNonMonetaryDealEndedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsTerminationTypeFuncInit(&this.Termination, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  if err := XsSubjectNonMonetaryObligationTypeFuncInit(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

func (this *XsNonMonetaryDealEndedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if  ctx.op == "D" && this.Termination.Same(ctx.personDeal.Termination)  {
    ctx.personDeal.Termination = nil
  }
  if  ctx.op != "D" &&  !this.Termination.Same(ctx.personDeal.Termination) {
    ctx.personDeal.Termination = &this.Termination
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  if   !this.SubjectNonMonetaryObligation.Same(ctx.personDeal.SubjectNonMonetaryObligation) {
    ctx.personDeal.SubjectNonMonetaryObligation = &this.SubjectNonMonetaryObligation
  }
  return nil
}

func (this *XsNonMonetaryDealEndedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsTerminationTypeFuncValidate(&this.Termination, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  if err := XsSubjectNonMonetaryObligationTypeFuncValidate(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

// 2.6 Изменились сведения о судебном споре или требовании по обязательству
type XsDealClaimChangedTypeFunc func(this *XsDealClaimChangedType, ctx *ParseContext) error

var XsDealClaimChangedTypeFuncInit XsDealClaimChangedTypeFunc = (*XsDealClaimChangedType).Init

var XsDealClaimChangedTypeFuncLoad XsDealClaimChangedTypeFunc = (*XsDealClaimChangedType).Load

var XsDealClaimChangedTypeFuncValidate XsDealClaimChangedTypeFunc = (*XsDealClaimChangedType).Validate

func (this *XsDealClaimChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  for _, v := range (*this).Lawsuit {
    if err := XsLawsuitTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsDealClaimChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  foundLawsuits := make(map[int]bool)
  for i, v := range ctx.personDeal.Lawsuit {
    for k, w := range this.Lawsuit {
      if w.Same(&v) {
        ctx.personDeal.Lawsuit[i] = w
        foundLawsuits[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Lawsuit {
      if !foundLawsuits[i] {
        ctx.personDeal.Lawsuit = append(ctx.personDeal.Lawsuit, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundLawsuits) > 0 {
      var newLawsuits []XsLawsuitType
      for i, v := range ctx.personDeal.Lawsuit {
        if !foundLawsuits[i] {
          newLawsuits = append(newLawsuits, v)
        }
      }
      ctx.personDeal.Lawsuit = newLawsuits
    }
  }
  return nil
}

func (this *XsDealClaimChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  for _, v := range (*this).Lawsuit {
    if err := XsLawsuitTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// 2.7 Квалифицированное бюро получило от бюро данные для формирования сведений о среднемесячных платежах субъекта
type XsAvgPaymentsSuppliedTypeFunc func(this *XsAvgPaymentsSuppliedType, ctx *ParseContext) error

var XsAvgPaymentsSuppliedTypeFuncInit XsAvgPaymentsSuppliedTypeFunc = (*XsAvgPaymentsSuppliedType).Init

var XsAvgPaymentsSuppliedTypeFuncLoad XsAvgPaymentsSuppliedTypeFunc = (*XsAvgPaymentsSuppliedType).Load

var XsAvgPaymentsSuppliedTypeFuncValidate XsAvgPaymentsSuppliedTypeFunc = (*XsAvgPaymentsSuppliedType).Validate

func (this *XsAvgPaymentsSuppliedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsAvgPaymentTypeFuncInit(&this.AvgPayments, ctx); err != nil { return err }
  return nil
}

func (this *XsAvgPaymentsSuppliedType) Load(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncLoad(&this.DealUid, ctx); err != nil { return err }
  if err := XsAvgPaymentTypeFuncLoad(&this.AvgPayments, ctx); err != nil { return err }
  return nil
}

func (this *XsAvgPaymentsSuppliedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsAvgPaymentTypeFuncValidate(&this.AvgPayments, ctx); err != nil { return err }
  return nil
}

// 2.8.1 Конкурсное производство в отношении источника открылось
type XsSourceBankruptcyStartedTypeFunc func(this *XsSourceBankruptcyStartedType, ctx *ParseContext) error

var XsSourceBankruptcyStartedTypeFuncInit XsSourceBankruptcyStartedTypeFunc = (*XsSourceBankruptcyStartedType).Init

var XsSourceBankruptcyStartedTypeFuncLoad XsSourceBankruptcyStartedTypeFunc = (*XsSourceBankruptcyStartedType).Load

var XsSourceBankruptcyStartedTypeFuncValidate XsSourceBankruptcyStartedTypeFunc = (*XsSourceBankruptcyStartedType).Validate

func (this *XsSourceBankruptcyStartedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsSourceBankruptcyTypeFuncInit(&this.SourceBankruptcy, ctx); err != nil { return err }
  return nil
}

func (this *XsSourceBankruptcyStartedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if   !this.Amount.Same(ctx.personDeal.Amount) {
    ctx.personDeal.Amount = &this.Amount
  }
  if  ctx.op == "D" && this.SourceBankruptcy.Same(ctx.personDeal.SourceBankruptcy)  {
    ctx.personDeal.SourceBankruptcy = nil
  }
  if  ctx.op != "D" &&  !this.SourceBankruptcy.Same(ctx.personDeal.SourceBankruptcy) {
    ctx.personDeal.SourceBankruptcy = &this.SourceBankruptcy
  }
  return nil
}

func (this *XsSourceBankruptcyStartedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsSourceBankruptcyTypeFuncValidate(&this.SourceBankruptcy, ctx); err != nil { return err }
  return nil
}

// 2.8.2 В ходе конкурсного производства изменились сведения об исполнении субъектом своего обязательства или его части либо обязательство прекратилось
type XsSourceBankruptcyChangedTypeFunc func(this *XsSourceBankruptcyChangedType, ctx *ParseContext) error

var XsSourceBankruptcyChangedTypeFuncInit XsSourceBankruptcyChangedTypeFunc = (*XsSourceBankruptcyChangedType).Init

var XsSourceBankruptcyChangedTypeFuncLoad XsSourceBankruptcyChangedTypeFunc = (*XsSourceBankruptcyChangedType).Load

var XsSourceBankruptcyChangedTypeFuncValidate XsSourceBankruptcyChangedTypeFunc = (*XsSourceBankruptcyChangedType).Validate

func (this *XsSourceBankruptcyChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceBankruptcyTypeFuncInit(&this.SourceBankruptcy, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsSourceBankruptcyChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  if   !this.SourceBankruptcy.Same(ctx.personDeal.SourceBankruptcy) {
    ctx.personDeal.SourceBankruptcy = &this.SourceBankruptcy
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  return nil
}

func (this *XsSourceBankruptcyChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceBankruptcyTypeFuncValidate(&this.SourceBankruptcy, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 2.8.3 Конкурсное производство в отношении источника завершилось
type XsSourceBankruptcyEndedTypeFunc func(this *XsSourceBankruptcyEndedType, ctx *ParseContext) error

var XsSourceBankruptcyEndedTypeFuncInit XsSourceBankruptcyEndedTypeFunc = (*XsSourceBankruptcyEndedType).Init

var XsSourceBankruptcyEndedTypeFuncLoad XsSourceBankruptcyEndedTypeFunc = (*XsSourceBankruptcyEndedType).Load

var XsSourceBankruptcyEndedTypeFuncValidate XsSourceBankruptcyEndedTypeFunc = (*XsSourceBankruptcyEndedType).Validate

func (this *XsSourceBankruptcyEndedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceBankruptcyTypeFuncInit(&this.SourceBankruptcy, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsSourceBankruptcyEndedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  if   !this.SourceBankruptcy.Same(ctx.personDeal.SourceBankruptcy) {
    ctx.personDeal.SourceBankruptcy = &this.SourceBankruptcy
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  return nil
}

func (this *XsSourceBankruptcyEndedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceBankruptcyTypeFuncValidate(&this.SourceBankruptcy, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 2.9.1 Процесс ликвидации источника начался
type XsSourceLiquidationStartedTypeFunc func(this *XsSourceLiquidationStartedType, ctx *ParseContext) error

var XsSourceLiquidationStartedTypeFuncInit XsSourceLiquidationStartedTypeFunc = (*XsSourceLiquidationStartedType).Init

var XsSourceLiquidationStartedTypeFuncLoad XsSourceLiquidationStartedTypeFunc = (*XsSourceLiquidationStartedType).Load

var XsSourceLiquidationStartedTypeFuncValidate XsSourceLiquidationStartedTypeFunc = (*XsSourceLiquidationStartedType).Validate

func (this *XsSourceLiquidationStartedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsSourceLiquidationTypeFuncInit(&this.SourceLiquidation, ctx); err != nil { return err }
  return nil
}

func (this *XsSourceLiquidationStartedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if   !this.Amount.Same(ctx.personDeal.Amount) {
    ctx.personDeal.Amount = &this.Amount
  }
  if  ctx.op == "D" && this.SourceLiquidation.Same(ctx.personDeal.SourceLiquidation)  {
    ctx.personDeal.SourceLiquidation = nil
  }
  if  ctx.op != "D" &&  !this.SourceLiquidation.Same(ctx.personDeal.SourceLiquidation) {
    ctx.personDeal.SourceLiquidation = &this.SourceLiquidation
  }
  return nil
}

func (this *XsSourceLiquidationStartedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsSourceLiquidationTypeFuncValidate(&this.SourceLiquidation, ctx); err != nil { return err }
  return nil
}

// 2.9.2 В ходе процесса ликвидации изменились сведения об исполнении субъектом своего обязательства или его части либо обязательство прекратилось
type XsSourceLiquidationChangedTypeFunc func(this *XsSourceLiquidationChangedType, ctx *ParseContext) error

var XsSourceLiquidationChangedTypeFuncInit XsSourceLiquidationChangedTypeFunc = (*XsSourceLiquidationChangedType).Init

var XsSourceLiquidationChangedTypeFuncLoad XsSourceLiquidationChangedTypeFunc = (*XsSourceLiquidationChangedType).Load

var XsSourceLiquidationChangedTypeFuncValidate XsSourceLiquidationChangedTypeFunc = (*XsSourceLiquidationChangedType).Validate

func (this *XsSourceLiquidationChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceLiquidationTypeFuncInit(&this.SourceLiquidation, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsSourceLiquidationChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  if   !this.SourceLiquidation.Same(ctx.personDeal.SourceLiquidation) {
    ctx.personDeal.SourceLiquidation = &this.SourceLiquidation
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  return nil
}

func (this *XsSourceLiquidationChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceLiquidationTypeFuncValidate(&this.SourceLiquidation, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 2.9.3 Процесс ликвидации источника завершился
type XsSourceLiquidationEndedTypeFunc func(this *XsSourceLiquidationEndedType, ctx *ParseContext) error

var XsSourceLiquidationEndedTypeFuncInit XsSourceLiquidationEndedTypeFunc = (*XsSourceLiquidationEndedType).Init

var XsSourceLiquidationEndedTypeFuncLoad XsSourceLiquidationEndedTypeFunc = (*XsSourceLiquidationEndedType).Load

var XsSourceLiquidationEndedTypeFuncValidate XsSourceLiquidationEndedTypeFunc = (*XsSourceLiquidationEndedType).Validate

func (this *XsSourceLiquidationEndedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceLiquidationTypeFuncInit(&this.SourceLiquidation, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsSourceLiquidationEndedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  if   !this.SourceLiquidation.Same(ctx.personDeal.SourceLiquidation) {
    ctx.personDeal.SourceLiquidation = &this.SourceLiquidation
  }
  foundParticipation := -1
  for i, v := range ctx.person.Participation {
    if this.Participation.Same(&v) {
      ctx.person.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.person.Participation = append(ctx.person.Participation, this.Participation)
  }
  return nil
}

func (this *XsSourceLiquidationEndedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceLiquidationTypeFuncValidate(&this.SourceLiquidation, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 2.10 Источник прекратил передачу информации по обязательству
type XsDealInfoStoppedTypeFunc func(this *XsDealInfoStoppedType, ctx *ParseContext) error

var XsDealInfoStoppedTypeFuncInit XsDealInfoStoppedTypeFunc = (*XsDealInfoStoppedType).Init

var XsDealInfoStoppedTypeFuncLoad XsDealInfoStoppedTypeFunc = (*XsDealInfoStoppedType).Load

var XsDealInfoStoppedTypeFuncValidate XsDealInfoStoppedTypeFunc = (*XsDealInfoStoppedType).Validate

func (this *XsDealInfoStoppedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsStopInfoTypeFuncInit(&this.StopInfo, ctx); err != nil { return err }
  return nil
}

func (this *XsDealInfoStoppedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  foundStopInfo := -1
  for i, v := range ctx.personDeal.StopInfo {
    if this.StopInfo.Same(&v) {
      ctx.personDeal.StopInfo[i] = this.StopInfo
      foundStopInfo = i
    }
  }
  if ctx.op != "D" && foundStopInfo == -1 {
    ctx.personDeal.StopInfo = append(ctx.personDeal.StopInfo, this.StopInfo)
  }
  if ctx.op == "D" && foundStopInfo != -1 {
    ctx.personDeal.StopInfo = append(ctx.personDeal.StopInfo[:foundStopInfo], ctx.personDeal.StopInfo[foundStopInfo+1:]...)
  }
  return nil
}

func (this *XsDealInfoStoppedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsStopInfoTypeFuncValidate(&this.StopInfo, ctx); err != nil { return err }
  return nil
}

// 2.11 Права кредитора по обязательству перешли к другому лицу
type XsCreditorChangedTypeFunc func(this *XsCreditorChangedType, ctx *ParseContext) error

var XsCreditorChangedTypeFuncInit XsCreditorChangedTypeFunc = (*XsCreditorChangedType).Init

var XsCreditorChangedTypeFuncLoad XsCreditorChangedTypeFunc = (*XsCreditorChangedType).Load

var XsCreditorChangedTypeFuncValidate XsCreditorChangedTypeFunc = (*XsCreditorChangedType).Validate

func (this *XsCreditorChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsStopInfoTypeFuncInit(&this.StopInfo, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncInit(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncInit(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsCreditorChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  foundStopInfo := -1
  for i, v := range ctx.personDeal.StopInfo {
    if this.StopInfo.Same(&v) {
      ctx.personDeal.StopInfo[i] = this.StopInfo
      foundStopInfo = i
    }
  }
  if ctx.op != "D" && foundStopInfo == -1 {
    ctx.personDeal.StopInfo = append(ctx.personDeal.StopInfo, this.StopInfo)
  }
  if   !this.Accounting.Same(ctx.personDeal.Accounting) {
    ctx.personDeal.Accounting = &this.Accounting
  }
  if  ctx.op == "D" && this.OrgAcquirerInfo.Same(ctx.personDeal.OrgAcquirerInfo)  {
    ctx.personDeal.OrgAcquirerInfo = nil
  }
  if  ctx.op != "D" &&  !this.OrgAcquirerInfo.Same(ctx.personDeal.OrgAcquirerInfo) {
    ctx.personDeal.OrgAcquirerInfo = this.OrgAcquirerInfo
  }
  if  ctx.op == "D" && this.PersonAcquirerInfo.Same(ctx.personDeal.PersonAcquirerInfo)  {
    ctx.personDeal.PersonAcquirerInfo = nil
  }
  if  ctx.op != "D" &&  !this.PersonAcquirerInfo.Same(ctx.personDeal.PersonAcquirerInfo) {
    ctx.personDeal.PersonAcquirerInfo = this.PersonAcquirerInfo
  }
  return nil
}

func (this *XsCreditorChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsStopInfoTypeFuncValidate(&this.StopInfo, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncValidate(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncValidate(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  return nil
}

// 2.12 Изменились сведения об обслуживающей организации (в частности, заключен, изменен или расторгнут договор обслуживания)
type XsServiceOrgChangedTypeFunc func(this *XsServiceOrgChangedType, ctx *ParseContext) error

var XsServiceOrgChangedTypeFuncInit XsServiceOrgChangedTypeFunc = (*XsServiceOrgChangedType).Init

var XsServiceOrgChangedTypeFuncLoad XsServiceOrgChangedTypeFunc = (*XsServiceOrgChangedType).Load

var XsServiceOrgChangedTypeFuncValidate XsServiceOrgChangedTypeFunc = (*XsServiceOrgChangedType).Validate

func (this *XsServiceOrgChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsServiceOrgTypeFuncInit(&this.ServiceOrg, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  return nil
}

func (this *XsServiceOrgChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  if  ctx.op == "D" && this.ServiceOrg.Same(ctx.personDeal.ServiceOrg)  {
    ctx.personDeal.ServiceOrg = nil
  }
  if  ctx.op != "D" &&  !this.ServiceOrg.Same(ctx.personDeal.ServiceOrg) {
    ctx.personDeal.ServiceOrg = &this.ServiceOrg
  }
  if   !this.Accounting.Same(ctx.personDeal.Accounting) {
    ctx.personDeal.Accounting = &this.Accounting
  }
  return nil
}

func (this *XsServiceOrgChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsServiceOrgTypeFuncValidate(&this.ServiceOrg, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  return nil
}

// 1.1 Субъект обратился к источнику с предложением совершить сделку
type XsOrgSubjectAppliedTypeFunc func(this *XsOrgSubjectAppliedType, ctx *ParseContext) error

var XsOrgSubjectAppliedTypeFuncInit XsOrgSubjectAppliedTypeFunc = (*XsOrgSubjectAppliedType).Init

var XsOrgSubjectAppliedTypeFuncLoad XsOrgSubjectAppliedTypeFunc = (*XsOrgSubjectAppliedType).Load

var XsOrgSubjectAppliedTypeFuncValidate XsOrgSubjectAppliedTypeFunc = (*XsOrgSubjectAppliedType).Validate

func (this *XsOrgSubjectAppliedType) Init(ctx *ParseContext) error {
  if err := XsApplicationTypeFuncInit(&this.Application, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgSubjectAppliedType) Load(ctx *ParseContext) error {
  foundApplication := -1
  for i, v := range ctx.org.Application {
    if this.Application.Same(&v) {
      ctx.org.Application[i] = this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.org.Application = append(ctx.org.Application, this.Application)
  }
  if ctx.op == "D" && foundApplication != -1 {
    ctx.org.Application = append(ctx.org.Application[:foundApplication], ctx.org.Application[foundApplication+1:]...)
  }
  return nil
}

func (this *XsOrgSubjectAppliedType) Validate(ctx *ParseContext) error {
  if err := XsApplicationTypeFuncValidate(&this.Application, ctx); err != nil { return err }
  return nil
}

// 1.2 Источник одобрил обращение субъекта (направил ему оферту) или изменились сведения об обращении
type XsOrgApplicationChangedTypeFunc func(this *XsOrgApplicationChangedType, ctx *ParseContext) error

var XsOrgApplicationChangedTypeFuncInit XsOrgApplicationChangedTypeFunc = (*XsOrgApplicationChangedType).Init

var XsOrgApplicationChangedTypeFuncLoad XsOrgApplicationChangedTypeFunc = (*XsOrgApplicationChangedType).Load

var XsOrgApplicationChangedTypeFuncValidate XsOrgApplicationChangedTypeFunc = (*XsOrgApplicationChangedType).Validate

func (this *XsOrgApplicationChangedType) Init(ctx *ParseContext) error {
  if err := XsApplicationTypeFuncInit(&this.Application, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgApplicationChangedType) Load(ctx *ParseContext) error {
  foundApplication := -1
  for i, v := range ctx.org.Application {
    if this.Application.Same(&v) {
      ctx.org.Application[i] = this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.org.Application = append(ctx.org.Application, this.Application)
  }
  return nil
}

func (this *XsOrgApplicationChangedType) Validate(ctx *ParseContext) error {
  if err := XsApplicationTypeFuncValidate(&this.Application, ctx); err != nil { return err }
  return nil
}

// 1.3 Источник отказался от совершения сделки по обращению субъекта
type XsOrgApplicationRejectedTypeFunc func(this *XsOrgApplicationRejectedType, ctx *ParseContext) error

var XsOrgApplicationRejectedTypeFuncInit XsOrgApplicationRejectedTypeFunc = (*XsOrgApplicationRejectedType).Init

var XsOrgApplicationRejectedTypeFuncLoad XsOrgApplicationRejectedTypeFunc = (*XsOrgApplicationRejectedType).Load

var XsOrgApplicationRejectedTypeFuncValidate XsOrgApplicationRejectedTypeFunc = (*XsOrgApplicationRejectedType).Validate

func (this *XsOrgApplicationRejectedType) Init(ctx *ParseContext) error {
  if err := XsApplicationTypeFuncInit(&this.Application, ctx); err != nil { return err }
  if err := XsAppRejectTypeFuncInit(&this.AppReject, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgApplicationRejectedType) Load(ctx *ParseContext) error {
  foundApplication := -1
  for i, v := range ctx.org.Application {
    if this.Application.Same(&v) {
      ctx.org.Application[i] = this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.org.Application = append(ctx.org.Application, this.Application)
  }
  foundAppReject := -1
  for i, v := range ctx.org.AppReject {
    if this.AppReject.Same(&v) {
      ctx.org.AppReject[i] = this.AppReject
      foundAppReject = i
    }
  }
  if ctx.op != "D" && foundAppReject == -1 {
    ctx.org.AppReject = append(ctx.org.AppReject, this.AppReject)
  }
  if ctx.op == "D" && foundAppReject != -1 {
    ctx.org.AppReject = append(ctx.org.AppReject[:foundAppReject], ctx.org.AppReject[foundAppReject+1:]...)
  }
  return nil
}

func (this *XsOrgApplicationRejectedType) Validate(ctx *ParseContext) error {
  if err := XsApplicationTypeFuncValidate(&this.Application, ctx); err != nil { return err }
  if err := XsAppRejectTypeFuncValidate(&this.AppReject, ctx); err != nil { return err }
  return nil
}

// 1.4 Субъект и источник совершили сделку, кроме договора лизинга и поручительства по лизингу для денежного обязательства субъекта
type XsOrgMonetaryDealMadeTypeFunc func(this *XsOrgMonetaryDealMadeType, ctx *ParseContext) error

var XsOrgMonetaryDealMadeTypeFuncInit XsOrgMonetaryDealMadeTypeFunc = (*XsOrgMonetaryDealMadeType).Init

var XsOrgMonetaryDealMadeTypeFuncLoad XsOrgMonetaryDealMadeTypeFunc = (*XsOrgMonetaryDealMadeType).Load

var XsOrgMonetaryDealMadeTypeFuncValidate XsOrgMonetaryDealMadeTypeFunc = (*XsOrgMonetaryDealMadeType).Validate

func (this *XsOrgMonetaryDealMadeType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsJointDebtorsTypeFuncInit(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncInit(this.Application, ctx); err != nil { return err }
  }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncInit(&this.PaymentTerms, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgMonetaryDealMadeType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  foundPrevOrgs := make(map[int]bool)
  for i, v := range ctx.org.PrevOrg {
    for k, w := range this.PrevOrg {
      if w.Same(&v) {
        ctx.org.PrevOrg[i] = w
        foundPrevOrgs[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevOrg {
      if !foundPrevOrgs[i] {
        ctx.org.PrevOrg = append(ctx.org.PrevOrg, w)
      }
    }
  }
  if   !this.JointDebtors.Same(ctx.orgDeal.JointDebtors) {
    ctx.orgDeal.JointDebtors = &this.JointDebtors
  }
  if   !this.Accounting.Same(ctx.orgDeal.Accounting) {
    ctx.orgDeal.Accounting = &this.Accounting
  }
  foundApplication := -1
  for i, v := range ctx.org.Application {
    if this.Application.Same(&v) {
      ctx.org.Application[i] = *this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.org.Application = append(ctx.org.Application, *this.Application)
  }
  if   !this.Amount.Same(ctx.orgDeal.Amount) {
    ctx.orgDeal.Amount = &this.Amount
  }
  if   !this.PaymentTerms.Same(ctx.orgDeal.PaymentTerms) {
    ctx.orgDeal.PaymentTerms = &this.PaymentTerms
  }
  return nil
}

func (this *XsOrgMonetaryDealMadeType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsJointDebtorsTypeFuncValidate(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncValidate(this.Application, ctx); err != nil { return err }
  }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncValidate(&this.PaymentTerms, ctx); err != nil { return err }
  return nil
}

// 1.4 Субъект и источник совершили сделку, кроме договора лизинга и поручительства по лизингу для неденежного обязательства источника
type XsOrgSourceNonMonetaryDealMadeTypeFunc func(this *XsOrgSourceNonMonetaryDealMadeType, ctx *ParseContext) error

var XsOrgSourceNonMonetaryDealMadeTypeFuncInit XsOrgSourceNonMonetaryDealMadeTypeFunc = (*XsOrgSourceNonMonetaryDealMadeType).Init

var XsOrgSourceNonMonetaryDealMadeTypeFuncLoad XsOrgSourceNonMonetaryDealMadeTypeFunc = (*XsOrgSourceNonMonetaryDealMadeType).Load

var XsOrgSourceNonMonetaryDealMadeTypeFuncValidate XsOrgSourceNonMonetaryDealMadeTypeFunc = (*XsOrgSourceNonMonetaryDealMadeType).Validate

func (this *XsOrgSourceNonMonetaryDealMadeType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsJointDebtorsTypeFuncInit(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncInit(this.Application, ctx); err != nil { return err }
  }
  if err := XsSourceNonMonetaryObligationTypeFuncInit(&this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgSourceNonMonetaryDealMadeType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  foundPrevOrgs := make(map[int]bool)
  for i, v := range ctx.org.PrevOrg {
    for k, w := range this.PrevOrg {
      if w.Same(&v) {
        ctx.org.PrevOrg[i] = w
        foundPrevOrgs[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevOrg {
      if !foundPrevOrgs[i] {
        ctx.org.PrevOrg = append(ctx.org.PrevOrg, w)
      }
    }
  }
  if   !this.JointDebtors.Same(ctx.orgDeal.JointDebtors) {
    ctx.orgDeal.JointDebtors = &this.JointDebtors
  }
  if   !this.Accounting.Same(ctx.orgDeal.Accounting) {
    ctx.orgDeal.Accounting = &this.Accounting
  }
  foundApplication := -1
  for i, v := range ctx.org.Application {
    if this.Application.Same(&v) {
      ctx.org.Application[i] = *this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.org.Application = append(ctx.org.Application, *this.Application)
  }
  if   !this.SourceNonMonetaryObligation.Same(ctx.orgDeal.SourceNonMonetaryObligation) {
    ctx.orgDeal.SourceNonMonetaryObligation = &this.SourceNonMonetaryObligation
  }
  return nil
}

func (this *XsOrgSourceNonMonetaryDealMadeType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsJointDebtorsTypeFuncValidate(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncValidate(this.Application, ctx); err != nil { return err }
  }
  if err := XsSourceNonMonetaryObligationTypeFuncValidate(&this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

// 1.4 Субъект и источник совершили сделку, кроме договора лизинга и поручительства по лизингу для неденежного обязательства субъекта  
type XsOrgNonMonetaryDealMadeTypeFunc func(this *XsOrgNonMonetaryDealMadeType, ctx *ParseContext) error

var XsOrgNonMonetaryDealMadeTypeFuncInit XsOrgNonMonetaryDealMadeTypeFunc = (*XsOrgNonMonetaryDealMadeType).Init

var XsOrgNonMonetaryDealMadeTypeFuncLoad XsOrgNonMonetaryDealMadeTypeFunc = (*XsOrgNonMonetaryDealMadeType).Load

var XsOrgNonMonetaryDealMadeTypeFuncValidate XsOrgNonMonetaryDealMadeTypeFunc = (*XsOrgNonMonetaryDealMadeType).Validate

func (this *XsOrgNonMonetaryDealMadeType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsJointDebtorsTypeFuncInit(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncInit(this.Application, ctx); err != nil { return err }
  }
  if err := XsSubjectNonMonetaryObligationTypeFuncInit(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgNonMonetaryDealMadeType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  foundPrevOrgs := make(map[int]bool)
  for i, v := range ctx.org.PrevOrg {
    for k, w := range this.PrevOrg {
      if w.Same(&v) {
        ctx.org.PrevOrg[i] = w
        foundPrevOrgs[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevOrg {
      if !foundPrevOrgs[i] {
        ctx.org.PrevOrg = append(ctx.org.PrevOrg, w)
      }
    }
  }
  if   !this.JointDebtors.Same(ctx.orgDeal.JointDebtors) {
    ctx.orgDeal.JointDebtors = &this.JointDebtors
  }
  if   !this.Accounting.Same(ctx.orgDeal.Accounting) {
    ctx.orgDeal.Accounting = &this.Accounting
  }
  foundApplication := -1
  for i, v := range ctx.org.Application {
    if this.Application.Same(&v) {
      ctx.org.Application[i] = *this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.org.Application = append(ctx.org.Application, *this.Application)
  }
  if   !this.SubjectNonMonetaryObligation.Same(ctx.orgDeal.SubjectNonMonetaryObligation) {
    ctx.orgDeal.SubjectNonMonetaryObligation = &this.SubjectNonMonetaryObligation
  }
  return nil
}

func (this *XsOrgNonMonetaryDealMadeType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsJointDebtorsTypeFuncValidate(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncValidate(this.Application, ctx); err != nil { return err }
  }
  if err := XsSubjectNonMonetaryObligationTypeFuncValidate(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

// 1.4.1 Субъект и источник заключили договор лизинга либо поручительства по лизингу и предмет лизинга передан лизингополучателю
type XsOrgLeasingDealMadeTypeFunc func(this *XsOrgLeasingDealMadeType, ctx *ParseContext) error

var XsOrgLeasingDealMadeTypeFuncInit XsOrgLeasingDealMadeTypeFunc = (*XsOrgLeasingDealMadeType).Init

var XsOrgLeasingDealMadeTypeFuncLoad XsOrgLeasingDealMadeTypeFunc = (*XsOrgLeasingDealMadeType).Load

var XsOrgLeasingDealMadeTypeFuncValidate XsOrgLeasingDealMadeTypeFunc = (*XsOrgLeasingDealMadeType).Validate

func (this *XsOrgLeasingDealMadeType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsJointDebtorsTypeFuncInit(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncInit(this.Application, ctx); err != nil { return err }
  }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncInit(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncInit(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncInit(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncInit(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncInit(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsSourceNonMonetaryObligationTypeFuncInit(&this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgLeasingDealMadeType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  foundPrevOrgs := make(map[int]bool)
  for i, v := range ctx.org.PrevOrg {
    for k, w := range this.PrevOrg {
      if w.Same(&v) {
        ctx.org.PrevOrg[i] = w
        foundPrevOrgs[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevOrg {
      if !foundPrevOrgs[i] {
        ctx.org.PrevOrg = append(ctx.org.PrevOrg, w)
      }
    }
  }
  if   !this.JointDebtors.Same(ctx.orgDeal.JointDebtors) {
    ctx.orgDeal.JointDebtors = &this.JointDebtors
  }
  if   !this.Accounting.Same(ctx.orgDeal.Accounting) {
    ctx.orgDeal.Accounting = &this.Accounting
  }
  foundApplication := -1
  for i, v := range ctx.org.Application {
    if this.Application.Same(&v) {
      ctx.org.Application[i] = *this.Application
      foundApplication = i
    }
  }
  if ctx.op != "D" && foundApplication == -1 {
    ctx.org.Application = append(ctx.org.Application, *this.Application)
  }
  if   !this.Amount.Same(ctx.orgDeal.Amount) {
    ctx.orgDeal.Amount = &this.Amount
  }
  if   !this.PaymentTerms.Same(ctx.orgDeal.PaymentTerms) {
    ctx.orgDeal.PaymentTerms = &this.PaymentTerms
  }
  if   !this.FundDate.Same(ctx.orgDeal.FundDate) {
    ctx.orgDeal.FundDate = &this.FundDate
  }
  foundArrear := -1
  for i, v := range ctx.orgDeal.Arrear {
    if this.Arrear.Same(&v) {
      ctx.orgDeal.Arrear[i] = this.Arrear
      foundArrear = i
    }
  }
  if ctx.op != "D" && foundArrear == -1 {
    ctx.orgDeal.Arrear = append(ctx.orgDeal.Arrear, this.Arrear)
  }
  foundDueArrear := -1
  for i, v := range ctx.orgDeal.DueArrear {
    if this.DueArrear.Same(&v) {
      ctx.orgDeal.DueArrear[i] = this.DueArrear
      foundDueArrear = i
    }
  }
  if ctx.op != "D" && foundDueArrear == -1 {
    ctx.orgDeal.DueArrear = append(ctx.orgDeal.DueArrear, this.DueArrear)
  }
  foundPastdueArrear := -1
  for i, v := range ctx.orgDeal.PastdueArrear {
    if this.PastdueArrear.Same(&v) {
      ctx.orgDeal.PastdueArrear[i] = this.PastdueArrear
      foundPastdueArrear = i
    }
  }
  if ctx.op != "D" && foundPastdueArrear == -1 {
    ctx.orgDeal.PastdueArrear = append(ctx.orgDeal.PastdueArrear, this.PastdueArrear)
  }
  foundPaymentss := make(map[int]bool)
  for i, v := range ctx.orgDeal.Payments {
    for k, w := range this.Payments {
      if w.Same(&v) {
        ctx.orgDeal.Payments[i] = w
        foundPaymentss[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Payments {
      if !foundPaymentss[i] {
        ctx.orgDeal.Payments = append(ctx.orgDeal.Payments, w)
      }
    }
  }
  if   !this.SourceNonMonetaryObligation.Same(ctx.orgDeal.SourceNonMonetaryObligation) {
    ctx.orgDeal.SourceNonMonetaryObligation = &this.SourceNonMonetaryObligation
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  return nil
}

func (this *XsOrgLeasingDealMadeType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsJointDebtorsTypeFuncValidate(&this.JointDebtors, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if this.Application != nil {
    if err := XsApplicationTypeFuncValidate(this.Application, ctx); err != nil { return err }
  }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncValidate(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncValidate(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncValidate(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncValidate(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncValidate(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsSourceNonMonetaryObligationTypeFuncValidate(&this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 1.5 Для принудительного исполнения передано требование о взыскании долга по алиментам, платы за жилое помещение, коммунальные услуги или услуги связи
type XsOrgDebtClaimSubmittedTypeFunc func(this *XsOrgDebtClaimSubmittedType, ctx *ParseContext) error

var XsOrgDebtClaimSubmittedTypeFuncInit XsOrgDebtClaimSubmittedTypeFunc = (*XsOrgDebtClaimSubmittedType).Init

var XsOrgDebtClaimSubmittedTypeFuncLoad XsOrgDebtClaimSubmittedTypeFunc = (*XsOrgDebtClaimSubmittedType).Load

var XsOrgDebtClaimSubmittedTypeFuncValidate XsOrgDebtClaimSubmittedTypeFunc = (*XsOrgDebtClaimSubmittedType).Validate

func (this *XsOrgDebtClaimSubmittedType) Init(ctx *ParseContext) error {
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsCollectionTypeFuncInit(&this.Collection, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgDebtClaimSubmittedType) Load(ctx *ParseContext) error {
  foundPrevOrgs := make(map[int]bool)
  for i, v := range ctx.org.PrevOrg {
    for k, w := range this.PrevOrg {
      if w.Same(&v) {
        ctx.org.PrevOrg[i] = w
        foundPrevOrgs[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevOrg {
      if !foundPrevOrgs[i] {
        ctx.org.PrevOrg = append(ctx.org.PrevOrg, w)
      }
    }
  }
  foundCollection := -1
  for i, v := range ctx.org.Collection {
    if this.Collection.Same(&v) {
      ctx.org.Collection[i] = this.Collection
      foundCollection = i
    }
  }
  if ctx.op != "D" && foundCollection == -1 {
    ctx.org.Collection = append(ctx.org.Collection, this.Collection)
  }
  if ctx.op == "D" && foundCollection != -1 {
    ctx.org.Collection = append(ctx.org.Collection[:foundCollection], ctx.org.Collection[foundCollection+1:]...)
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  return nil
}

func (this *XsOrgDebtClaimSubmittedType) Validate(ctx *ParseContext) error {
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsCollectionTypeFuncValidate(&this.Collection, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 1.6 Изменились сведения о взыскании долга по алиментам, платы за жилое помещение, коммунальные услуги или услуги связи либо указанное требование погашено
type XsOrgDebtClaimChangedTypeFunc func(this *XsOrgDebtClaimChangedType, ctx *ParseContext) error

var XsOrgDebtClaimChangedTypeFuncInit XsOrgDebtClaimChangedTypeFunc = (*XsOrgDebtClaimChangedType).Init

var XsOrgDebtClaimChangedTypeFuncLoad XsOrgDebtClaimChangedTypeFunc = (*XsOrgDebtClaimChangedType).Load

var XsOrgDebtClaimChangedTypeFuncValidate XsOrgDebtClaimChangedTypeFunc = (*XsOrgDebtClaimChangedType).Validate

func (this *XsOrgDebtClaimChangedType) Init(ctx *ParseContext) error {
  if err := XsCollectionTypeFuncInit(&this.Collection, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgDebtClaimChangedType) Load(ctx *ParseContext) error {
  foundCollection := -1
  for i, v := range ctx.org.Collection {
    if this.Collection.Same(&v) {
      ctx.org.Collection[i] = this.Collection
      foundCollection = i
    }
  }
  if ctx.op != "D" && foundCollection == -1 {
    ctx.org.Collection = append(ctx.org.Collection, this.Collection)
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  return nil
}

func (this *XsOrgDebtClaimChangedType) Validate(ctx *ParseContext) error {
  if err := XsCollectionTypeFuncValidate(&this.Collection, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 1.7 и 1.8 Изменились сведения титульной части кредитной истории субъекта
type XsOrgSubjectMainChangedTypeFunc func(this *XsOrgSubjectMainChangedType, ctx *ParseContext) error

var XsOrgSubjectMainChangedTypeFuncInit XsOrgSubjectMainChangedTypeFunc = (*XsOrgSubjectMainChangedType).Init

var XsOrgSubjectMainChangedTypeFuncLoad XsOrgSubjectMainChangedTypeFunc = (*XsOrgSubjectMainChangedType).Load

var XsOrgSubjectMainChangedTypeFuncValidate XsOrgSubjectMainChangedTypeFunc = (*XsOrgSubjectMainChangedType).Validate

func (this *XsOrgSubjectMainChangedType) Init(ctx *ParseContext) error {
  if err := XsOrgNameTypeFuncInit(&this.Name, ctx); err != nil { return err }
  if err := XsOrgAddressTypeFuncInit(&this.Address, ctx); err != nil { return err }
  if err := XsOrgRegNumTypeFuncInit(&this.RegNum, ctx); err != nil { return err }
  if this.Tax != nil {
    if err := XsOrgTaxTypeFuncInit(this.Tax, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgChange {
    if err := XsOrgChangeTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsOrgSubjectMainChangedType) Load(ctx *ParseContext) error {
  if   !this.Name.Same(&ctx.org.Name) {
    ctx.org.Name = this.Name
  }
  if   !this.Address.Same(&ctx.org.Address) {
    ctx.org.Address = this.Address
  }
  if   !this.RegNum.Same(&ctx.org.RegNum) {
    ctx.org.RegNum = this.RegNum
  }
  if  ctx.op == "D" && this.Tax.Same(ctx.org.Tax)  {
    ctx.org.Tax = nil
  }
  if  ctx.op != "D" &&  !this.Tax.Same(ctx.org.Tax) {
    ctx.org.Tax = this.Tax
  }
  foundOrgChanges := make(map[int]bool)
  for i, v := range ctx.org.OrgChange {
    for k, w := range this.OrgChange {
      if w.Same(&v) {
        ctx.org.OrgChange[i] = w
        foundOrgChanges[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.OrgChange {
      if !foundOrgChanges[i] {
        ctx.org.OrgChange = append(ctx.org.OrgChange, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundOrgChanges) > 0 {
      var newOrgChanges []XsOrgChangeType
      for i, v := range ctx.org.OrgChange {
        if !foundOrgChanges[i] {
          newOrgChanges = append(newOrgChanges, v)
        }
      }
      ctx.org.OrgChange = newOrgChanges
    }
  }
  return nil
}

func (this *XsOrgSubjectMainChangedType) Validate(ctx *ParseContext) error {
  if err := XsOrgNameTypeFuncValidate(&this.Name, ctx); err != nil { return err }
  if err := XsOrgAddressTypeFuncValidate(&this.Address, ctx); err != nil { return err }
  if err := XsOrgRegNumTypeFuncValidate(&this.RegNum, ctx); err != nil { return err }
  if this.Tax != nil {
    if err := XsOrgTaxTypeFuncValidate(this.Tax, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgChange {
    if err := XsOrgChangeTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// 1.9 Изменились сведения о субъекте в основной части кредитной истории, кроме сведений о дееспособности, банкротстве, индивидуальном рейтинге и кредитной оценке
type XsOrgSubjectSpecialChangedTypeFunc func(this *XsOrgSubjectSpecialChangedType, ctx *ParseContext) error

var XsOrgSubjectSpecialChangedTypeFuncInit XsOrgSubjectSpecialChangedTypeFunc = (*XsOrgSubjectSpecialChangedType).Init

var XsOrgSubjectSpecialChangedTypeFuncLoad XsOrgSubjectSpecialChangedTypeFunc = (*XsOrgSubjectSpecialChangedType).Load

var XsOrgSubjectSpecialChangedTypeFuncValidate XsOrgSubjectSpecialChangedTypeFunc = (*XsOrgSubjectSpecialChangedType).Validate

func (this *XsOrgSubjectSpecialChangedType) Init(ctx *ParseContext) error {
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsOrgSubjectSpecialChangedType) Load(ctx *ParseContext) error {
  foundPrevOrgs := make(map[int]bool)
  for i, v := range ctx.org.PrevOrg {
    for k, w := range this.PrevOrg {
      if w.Same(&v) {
        ctx.org.PrevOrg[i] = w
        foundPrevOrgs[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevOrg {
      if !foundPrevOrgs[i] {
        ctx.org.PrevOrg = append(ctx.org.PrevOrg, w)
      }
    }
  }
  return nil
}

func (this *XsOrgSubjectSpecialChangedType) Validate(ctx *ParseContext) error {
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// 1.10 и 1.11 Изменились сведения о дееспособности субъекта
type XsOrgSubjectCapacityChangedTypeFunc func(this *XsOrgSubjectCapacityChangedType, ctx *ParseContext) error

var XsOrgSubjectCapacityChangedTypeFuncInit XsOrgSubjectCapacityChangedTypeFunc = (*XsOrgSubjectCapacityChangedType).Init

var XsOrgSubjectCapacityChangedTypeFuncLoad XsOrgSubjectCapacityChangedTypeFunc = (*XsOrgSubjectCapacityChangedType).Load

var XsOrgSubjectCapacityChangedTypeFuncValidate XsOrgSubjectCapacityChangedTypeFunc = (*XsOrgSubjectCapacityChangedType).Validate

func (this *XsOrgSubjectCapacityChangedType) Init(ctx *ParseContext) error {
  if err := XsLegalCapacityTypeFuncInit(&this.LegalCapacity, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgSubjectCapacityChangedType) Load(ctx *ParseContext) error {
  if err := XsLegalCapacityTypeFuncLoad(&this.LegalCapacity, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgSubjectCapacityChangedType) Validate(ctx *ParseContext) error {
  if err := XsLegalCapacityTypeFuncValidate(&this.LegalCapacity, ctx); err != nil { return err }
  return nil
}

// 1.12 Изменились сведения по делу о банкротстве субъекта
type XsOrgSubjectBankruptcyChangedTypeFunc func(this *XsOrgSubjectBankruptcyChangedType, ctx *ParseContext) error

var XsOrgSubjectBankruptcyChangedTypeFuncInit XsOrgSubjectBankruptcyChangedTypeFunc = (*XsOrgSubjectBankruptcyChangedType).Init

var XsOrgSubjectBankruptcyChangedTypeFuncLoad XsOrgSubjectBankruptcyChangedTypeFunc = (*XsOrgSubjectBankruptcyChangedType).Load

var XsOrgSubjectBankruptcyChangedTypeFuncValidate XsOrgSubjectBankruptcyChangedTypeFunc = (*XsOrgSubjectBankruptcyChangedType).Validate

func (this *XsOrgSubjectBankruptcyChangedType) Init(ctx *ParseContext) error {
  if err := XsBankruptcyTypeFuncInit(&this.Bankruptcy, ctx); err != nil { return err }
  if err := XsOrgFulfillmentTypeFuncInit(&this.Fulfillment, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgSubjectBankruptcyChangedType) Load(ctx *ParseContext) error {
  foundBankruptcy := -1
  for i, v := range ctx.org.Bankruptcy {
    if this.Bankruptcy.Same(&v) {
      ctx.org.Bankruptcy[i] = this.Bankruptcy
      foundBankruptcy = i
    }
  }
  if ctx.op != "D" && foundBankruptcy == -1 {
    ctx.org.Bankruptcy = append(ctx.org.Bankruptcy, this.Bankruptcy)
  }
  foundFulfillment := -1
  for i, v := range ctx.org.Fulfillment {
    if this.Fulfillment.Same(&v) {
      ctx.org.Fulfillment[i] = this.Fulfillment
      foundFulfillment = i
    }
  }
  if ctx.op != "D" && foundFulfillment == -1 {
    ctx.org.Fulfillment = append(ctx.org.Fulfillment, this.Fulfillment)
  }
  if ctx.op == "D" && foundFulfillment != -1 {
    ctx.org.Fulfillment = append(ctx.org.Fulfillment[:foundFulfillment], ctx.org.Fulfillment[foundFulfillment+1:]...)
  }
  return nil
}

func (this *XsOrgSubjectBankruptcyChangedType) Validate(ctx *ParseContext) error {
  if err := XsBankruptcyTypeFuncValidate(&this.Bankruptcy, ctx); err != nil { return err }
  if err := XsOrgFulfillmentTypeFuncValidate(&this.Fulfillment, ctx); err != nil { return err }
  return nil
}

// 2.1 Изменились сведения об условиях обязательства субъекта для денежного обязательства
type XsOrgMonetaryDealChangedTypeFunc func(this *XsOrgMonetaryDealChangedType, ctx *ParseContext) error

var XsOrgMonetaryDealChangedTypeFuncInit XsOrgMonetaryDealChangedTypeFunc = (*XsOrgMonetaryDealChangedType).Init

var XsOrgMonetaryDealChangedTypeFuncLoad XsOrgMonetaryDealChangedTypeFunc = (*XsOrgMonetaryDealChangedType).Load

var XsOrgMonetaryDealChangedTypeFuncValidate XsOrgMonetaryDealChangedTypeFunc = (*XsOrgMonetaryDealChangedType).Validate

func (this *XsOrgMonetaryDealChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncInit(&this.JointDebtors, ctx); err != nil { return err }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncInit(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncInit(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncInit(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncInit(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncInit(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsOrgMonetaryDealChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  if  ctx.op == "D" && this.JointDebtors.Same(ctx.orgDeal.JointDebtors)  {
    ctx.orgDeal.JointDebtors = nil
  }
  if  ctx.op != "D" &&  !this.JointDebtors.Same(ctx.orgDeal.JointDebtors) {
    ctx.orgDeal.JointDebtors = &this.JointDebtors
  }
  foundChanges := make(map[int]bool)
  for i, v := range ctx.orgDeal.Change {
    for k, w := range this.Change {
      if w.Same(&v) {
        ctx.orgDeal.Change[i] = w
        foundChanges[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Change {
      if !foundChanges[i] {
        ctx.orgDeal.Change = append(ctx.orgDeal.Change, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundChanges) > 0 {
      var newChanges []XsDealChangeType
      for i, v := range ctx.orgDeal.Change {
        if !foundChanges[i] {
          newChanges = append(newChanges, v)
        }
      }
      ctx.orgDeal.Change = newChanges
    }
  }
  if  ctx.op == "D" && this.Accounting.Same(ctx.orgDeal.Accounting)  {
    ctx.orgDeal.Accounting = nil
  }
  if  ctx.op != "D" &&  !this.Accounting.Same(ctx.orgDeal.Accounting) {
    ctx.orgDeal.Accounting = &this.Accounting
  }
  if   !this.Amount.Same(ctx.orgDeal.Amount) {
    ctx.orgDeal.Amount = &this.Amount
  }
  if   !this.PaymentTerms.Same(ctx.orgDeal.PaymentTerms) {
    ctx.orgDeal.PaymentTerms = &this.PaymentTerms
  }
  if   !this.FundDate.Same(ctx.orgDeal.FundDate) {
    ctx.orgDeal.FundDate = &this.FundDate
  }
  foundArrear := -1
  for i, v := range ctx.orgDeal.Arrear {
    if this.Arrear.Same(&v) {
      ctx.orgDeal.Arrear[i] = this.Arrear
      foundArrear = i
    }
  }
  if ctx.op != "D" && foundArrear == -1 {
    ctx.orgDeal.Arrear = append(ctx.orgDeal.Arrear, this.Arrear)
  }
  if ctx.op == "D" && foundArrear != -1 {
    ctx.orgDeal.Arrear = append(ctx.orgDeal.Arrear[:foundArrear], ctx.orgDeal.Arrear[foundArrear+1:]...)
  }
  foundDueArrear := -1
  for i, v := range ctx.orgDeal.DueArrear {
    if this.DueArrear.Same(&v) {
      ctx.orgDeal.DueArrear[i] = this.DueArrear
      foundDueArrear = i
    }
  }
  if ctx.op != "D" && foundDueArrear == -1 {
    ctx.orgDeal.DueArrear = append(ctx.orgDeal.DueArrear, this.DueArrear)
  }
  if ctx.op == "D" && foundDueArrear != -1 {
    ctx.orgDeal.DueArrear = append(ctx.orgDeal.DueArrear[:foundDueArrear], ctx.orgDeal.DueArrear[foundDueArrear+1:]...)
  }
  foundPastdueArrear := -1
  for i, v := range ctx.orgDeal.PastdueArrear {
    if this.PastdueArrear.Same(&v) {
      ctx.orgDeal.PastdueArrear[i] = this.PastdueArrear
      foundPastdueArrear = i
    }
  }
  if ctx.op != "D" && foundPastdueArrear == -1 {
    ctx.orgDeal.PastdueArrear = append(ctx.orgDeal.PastdueArrear, this.PastdueArrear)
  }
  foundPaymentss := make(map[int]bool)
  for i, v := range ctx.orgDeal.Payments {
    for k, w := range this.Payments {
      if w.Same(&v) {
        ctx.orgDeal.Payments[i] = w
        foundPaymentss[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Payments {
      if !foundPaymentss[i] {
        ctx.orgDeal.Payments = append(ctx.orgDeal.Payments, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundPaymentss) > 0 {
      var newPaymentss []XsPaymentsType
      for i, v := range ctx.orgDeal.Payments {
        if !foundPaymentss[i] {
          newPaymentss = append(newPaymentss, v)
        }
      }
      ctx.orgDeal.Payments = newPaymentss
    }
  }
  return nil
}

func (this *XsOrgMonetaryDealChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncValidate(&this.JointDebtors, ctx); err != nil { return err }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncValidate(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncValidate(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncValidate(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncValidate(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncValidate(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// 2.1 Изменились сведения об условиях обязательства субъекта для неденежного обязательства
type XsOrgNonMonetaryDealChangedTypeFunc func(this *XsOrgNonMonetaryDealChangedType, ctx *ParseContext) error

var XsOrgNonMonetaryDealChangedTypeFuncInit XsOrgNonMonetaryDealChangedTypeFunc = (*XsOrgNonMonetaryDealChangedType).Init

var XsOrgNonMonetaryDealChangedTypeFuncLoad XsOrgNonMonetaryDealChangedTypeFunc = (*XsOrgNonMonetaryDealChangedType).Load

var XsOrgNonMonetaryDealChangedTypeFuncValidate XsOrgNonMonetaryDealChangedTypeFunc = (*XsOrgNonMonetaryDealChangedType).Validate

func (this *XsOrgNonMonetaryDealChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncInit(&this.JointDebtors, ctx); err != nil { return err }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if err := XsSubjectNonMonetaryObligationTypeFuncInit(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgNonMonetaryDealChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  if  ctx.op == "D" && this.JointDebtors.Same(ctx.orgDeal.JointDebtors)  {
    ctx.orgDeal.JointDebtors = nil
  }
  if  ctx.op != "D" &&  !this.JointDebtors.Same(ctx.orgDeal.JointDebtors) {
    ctx.orgDeal.JointDebtors = &this.JointDebtors
  }
  foundChanges := make(map[int]bool)
  for i, v := range ctx.orgDeal.Change {
    for k, w := range this.Change {
      if w.Same(&v) {
        ctx.orgDeal.Change[i] = w
        foundChanges[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Change {
      if !foundChanges[i] {
        ctx.orgDeal.Change = append(ctx.orgDeal.Change, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundChanges) > 0 {
      var newChanges []XsDealChangeType
      for i, v := range ctx.orgDeal.Change {
        if !foundChanges[i] {
          newChanges = append(newChanges, v)
        }
      }
      ctx.orgDeal.Change = newChanges
    }
  }
  if  ctx.op == "D" && this.Accounting.Same(ctx.orgDeal.Accounting)  {
    ctx.orgDeal.Accounting = nil
  }
  if  ctx.op != "D" &&  !this.Accounting.Same(ctx.orgDeal.Accounting) {
    ctx.orgDeal.Accounting = &this.Accounting
  }
  if   !this.SubjectNonMonetaryObligation.Same(ctx.orgDeal.SubjectNonMonetaryObligation) {
    ctx.orgDeal.SubjectNonMonetaryObligation = &this.SubjectNonMonetaryObligation
  }
  return nil
}

func (this *XsOrgNonMonetaryDealChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsJointDebtorsTypeFuncValidate(&this.JointDebtors, ctx); err != nil { return err }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if err := XsSubjectNonMonetaryObligationTypeFuncValidate(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

// 2.2 Субъекту передана сумма займа (кредита) либо субъект стал принципалом по гарантии или поручителем по сделке, кроме договора лизинга для денежного обязательства
type XsOrgMonetaryDealPrincipalChangedTypeFunc func(this *XsOrgMonetaryDealPrincipalChangedType, ctx *ParseContext) error

var XsOrgMonetaryDealPrincipalChangedTypeFuncInit XsOrgMonetaryDealPrincipalChangedTypeFunc = (*XsOrgMonetaryDealPrincipalChangedType).Init

var XsOrgMonetaryDealPrincipalChangedTypeFuncLoad XsOrgMonetaryDealPrincipalChangedTypeFunc = (*XsOrgMonetaryDealPrincipalChangedType).Load

var XsOrgMonetaryDealPrincipalChangedTypeFuncValidate XsOrgMonetaryDealPrincipalChangedTypeFunc = (*XsOrgMonetaryDealPrincipalChangedType).Validate

func (this *XsOrgMonetaryDealPrincipalChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncInit(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncInit(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncInit(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncInit(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncInit(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsOrgMonetaryDealPrincipalChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  if   !this.Amount.Same(ctx.orgDeal.Amount) {
    ctx.orgDeal.Amount = &this.Amount
  }
  if   !this.PaymentTerms.Same(ctx.orgDeal.PaymentTerms) {
    ctx.orgDeal.PaymentTerms = &this.PaymentTerms
  }
  if   !this.FundDate.Same(ctx.orgDeal.FundDate) {
    ctx.orgDeal.FundDate = &this.FundDate
  }
  foundArrear := -1
  for i, v := range ctx.orgDeal.Arrear {
    if this.Arrear.Same(&v) {
      ctx.orgDeal.Arrear[i] = this.Arrear
      foundArrear = i
    }
  }
  if ctx.op != "D" && foundArrear == -1 {
    ctx.orgDeal.Arrear = append(ctx.orgDeal.Arrear, this.Arrear)
  }
  foundDueArrear := -1
  for i, v := range ctx.orgDeal.DueArrear {
    if this.DueArrear.Same(&v) {
      ctx.orgDeal.DueArrear[i] = this.DueArrear
      foundDueArrear = i
    }
  }
  if ctx.op != "D" && foundDueArrear == -1 {
    ctx.orgDeal.DueArrear = append(ctx.orgDeal.DueArrear, this.DueArrear)
  }
  foundPastdueArrear := -1
  for i, v := range ctx.orgDeal.PastdueArrear {
    if this.PastdueArrear.Same(&v) {
      ctx.orgDeal.PastdueArrear[i] = this.PastdueArrear
      foundPastdueArrear = i
    }
  }
  if ctx.op != "D" && foundPastdueArrear == -1 {
    ctx.orgDeal.PastdueArrear = append(ctx.orgDeal.PastdueArrear, this.PastdueArrear)
  }
  foundPaymentss := make(map[int]bool)
  for i, v := range ctx.orgDeal.Payments {
    for k, w := range this.Payments {
      if w.Same(&v) {
        ctx.orgDeal.Payments[i] = w
        foundPaymentss[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Payments {
      if !foundPaymentss[i] {
        ctx.orgDeal.Payments = append(ctx.orgDeal.Payments, w)
      }
    }
  }
  return nil
}

func (this *XsOrgMonetaryDealPrincipalChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncValidate(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncValidate(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncValidate(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncValidate(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncValidate(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// 2.2 Субъекту передана сумма займа (кредита) либо субъект стал принципалом по гарантии или поручителем по сделке, кроме договора лизинга для денежного обязательства
type XsOrgNonMonetaryDealPrincipalChangedTypeFunc func(this *XsOrgNonMonetaryDealPrincipalChangedType, ctx *ParseContext) error

var XsOrgNonMonetaryDealPrincipalChangedTypeFuncInit XsOrgNonMonetaryDealPrincipalChangedTypeFunc = (*XsOrgNonMonetaryDealPrincipalChangedType).Init

var XsOrgNonMonetaryDealPrincipalChangedTypeFuncLoad XsOrgNonMonetaryDealPrincipalChangedTypeFunc = (*XsOrgNonMonetaryDealPrincipalChangedType).Load

var XsOrgNonMonetaryDealPrincipalChangedTypeFuncValidate XsOrgNonMonetaryDealPrincipalChangedTypeFunc = (*XsOrgNonMonetaryDealPrincipalChangedType).Validate

func (this *XsOrgNonMonetaryDealPrincipalChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  if err := XsSourceNonMonetaryObligationTypeFuncInit(&this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgNonMonetaryDealPrincipalChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  if   !this.SourceNonMonetaryObligation.Same(ctx.orgDeal.SourceNonMonetaryObligation) {
    ctx.orgDeal.SourceNonMonetaryObligation = &this.SourceNonMonetaryObligation
  }
  return nil
}

func (this *XsOrgNonMonetaryDealPrincipalChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  if err := XsSourceNonMonetaryObligationTypeFuncValidate(&this.SourceNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

// 2.3 Изменились сведения об исполнении обязательства субъектом, наступила ответственность поручителя или обязательство принципала возместить выплаченную сумму для денежного обязательства
type XsOrgMonetaryDealParticipationChangedTypeFunc func(this *XsOrgMonetaryDealParticipationChangedType, ctx *ParseContext) error

var XsOrgMonetaryDealParticipationChangedTypeFuncInit XsOrgMonetaryDealParticipationChangedTypeFunc = (*XsOrgMonetaryDealParticipationChangedType).Init

var XsOrgMonetaryDealParticipationChangedTypeFuncLoad XsOrgMonetaryDealParticipationChangedTypeFunc = (*XsOrgMonetaryDealParticipationChangedType).Load

var XsOrgMonetaryDealParticipationChangedTypeFuncValidate XsOrgMonetaryDealParticipationChangedTypeFunc = (*XsOrgMonetaryDealParticipationChangedType).Validate

func (this *XsOrgMonetaryDealParticipationChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncInit(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncInit(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncInit(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncInit(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncInit(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsOrgMonetaryDealParticipationChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  if   !this.Accounting.Same(ctx.orgDeal.Accounting) {
    ctx.orgDeal.Accounting = &this.Accounting
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  if ctx.op == "D" && foundParticipation != -1 {
    ctx.org.Participation = append(ctx.org.Participation[:foundParticipation], ctx.org.Participation[foundParticipation+1:]...)
  }
  if   !this.Amount.Same(ctx.orgDeal.Amount) {
    ctx.orgDeal.Amount = &this.Amount
  }
  if   !this.PaymentTerms.Same(ctx.orgDeal.PaymentTerms) {
    ctx.orgDeal.PaymentTerms = &this.PaymentTerms
  }
  if   !this.FundDate.Same(ctx.orgDeal.FundDate) {
    ctx.orgDeal.FundDate = &this.FundDate
  }
  foundArrear := -1
  for i, v := range ctx.orgDeal.Arrear {
    if this.Arrear.Same(&v) {
      ctx.orgDeal.Arrear[i] = this.Arrear
      foundArrear = i
    }
  }
  if ctx.op != "D" && foundArrear == -1 {
    ctx.orgDeal.Arrear = append(ctx.orgDeal.Arrear, this.Arrear)
  }
  foundDueArrear := -1
  for i, v := range ctx.orgDeal.DueArrear {
    if this.DueArrear.Same(&v) {
      ctx.orgDeal.DueArrear[i] = this.DueArrear
      foundDueArrear = i
    }
  }
  if ctx.op != "D" && foundDueArrear == -1 {
    ctx.orgDeal.DueArrear = append(ctx.orgDeal.DueArrear, this.DueArrear)
  }
  foundPastdueArrear := -1
  for i, v := range ctx.orgDeal.PastdueArrear {
    if this.PastdueArrear.Same(&v) {
      ctx.orgDeal.PastdueArrear[i] = this.PastdueArrear
      foundPastdueArrear = i
    }
  }
  if ctx.op != "D" && foundPastdueArrear == -1 {
    ctx.orgDeal.PastdueArrear = append(ctx.orgDeal.PastdueArrear, this.PastdueArrear)
  }
  foundPaymentss := make(map[int]bool)
  for i, v := range ctx.orgDeal.Payments {
    for k, w := range this.Payments {
      if w.Same(&v) {
        ctx.orgDeal.Payments[i] = w
        foundPaymentss[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Payments {
      if !foundPaymentss[i] {
        ctx.orgDeal.Payments = append(ctx.orgDeal.Payments, w)
      }
    }
  }
  return nil
}

func (this *XsOrgMonetaryDealParticipationChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncValidate(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncValidate(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncValidate(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncValidate(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncValidate(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// 2.3 Изменились сведения об исполнении обязательства субъектом, наступила ответственность поручителя или обязательство принципала возместить выплаченную сумму для неденежного обязательства
type XsOrgNonMonetaryDealParticipationChangedTypeFunc func(this *XsOrgNonMonetaryDealParticipationChangedType, ctx *ParseContext) error

var XsOrgNonMonetaryDealParticipationChangedTypeFuncInit XsOrgNonMonetaryDealParticipationChangedTypeFunc = (*XsOrgNonMonetaryDealParticipationChangedType).Init

var XsOrgNonMonetaryDealParticipationChangedTypeFuncLoad XsOrgNonMonetaryDealParticipationChangedTypeFunc = (*XsOrgNonMonetaryDealParticipationChangedType).Load

var XsOrgNonMonetaryDealParticipationChangedTypeFuncValidate XsOrgNonMonetaryDealParticipationChangedTypeFunc = (*XsOrgNonMonetaryDealParticipationChangedType).Validate

func (this *XsOrgNonMonetaryDealParticipationChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  if err := XsSubjectNonMonetaryObligationTypeFuncInit(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgNonMonetaryDealParticipationChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  if   !this.Accounting.Same(ctx.orgDeal.Accounting) {
    ctx.orgDeal.Accounting = &this.Accounting
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  if ctx.op == "D" && foundParticipation != -1 {
    ctx.org.Participation = append(ctx.org.Participation[:foundParticipation], ctx.org.Participation[foundParticipation+1:]...)
  }
  if   !this.SubjectNonMonetaryObligation.Same(ctx.orgDeal.SubjectNonMonetaryObligation) {
    ctx.orgDeal.SubjectNonMonetaryObligation = &this.SubjectNonMonetaryObligation
  }
  return nil
}

func (this *XsOrgNonMonetaryDealParticipationChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  if err := XsSubjectNonMonetaryObligationTypeFuncValidate(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

// 2.4 Изменились сведения об обеспечении исполнения обязательства
type XsOrgDealSecuringChangedTypeFunc func(this *XsOrgDealSecuringChangedType, ctx *ParseContext) error

var XsOrgDealSecuringChangedTypeFuncInit XsOrgDealSecuringChangedTypeFunc = (*XsOrgDealSecuringChangedType).Init

var XsOrgDealSecuringChangedTypeFuncLoad XsOrgDealSecuringChangedTypeFunc = (*XsOrgDealSecuringChangedType).Load

var XsOrgDealSecuringChangedTypeFuncValidate XsOrgDealSecuringChangedTypeFunc = (*XsOrgDealSecuringChangedType).Validate

func (this *XsOrgDealSecuringChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
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
  return nil
}

func (this *XsOrgDealSecuringChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  foundCollaterals := make(map[int]bool)
  for i, v := range ctx.orgDeal.Collateral {
    for k, w := range this.Collateral {
      if w.Same(&v) {
        ctx.orgDeal.Collateral[i] = w
        foundCollaterals[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Collateral {
      if !foundCollaterals[i] {
        ctx.orgDeal.Collateral = append(ctx.orgDeal.Collateral, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundCollaterals) > 0 {
      var newCollaterals []XsCollateralType
      for i, v := range ctx.orgDeal.Collateral {
        if !foundCollaterals[i] {
          newCollaterals = append(newCollaterals, v)
        }
      }
      ctx.orgDeal.Collateral = newCollaterals
    }
  }
  foundWarrantys := make(map[int]bool)
  for i, v := range ctx.orgDeal.Warranty {
    for k, w := range this.Warranty {
      if w.Same(&v) {
        ctx.orgDeal.Warranty[i] = w
        foundWarrantys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Warranty {
      if !foundWarrantys[i] {
        ctx.orgDeal.Warranty = append(ctx.orgDeal.Warranty, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundWarrantys) > 0 {
      var newWarrantys []XsWarrantyType
      for i, v := range ctx.orgDeal.Warranty {
        if !foundWarrantys[i] {
          newWarrantys = append(newWarrantys, v)
        }
      }
      ctx.orgDeal.Warranty = newWarrantys
    }
  }
  foundIndepGuarantees := make(map[int]bool)
  for i, v := range ctx.orgDeal.IndepGuarantee {
    for k, w := range this.IndepGuarantee {
      if w.Same(&v) {
        ctx.orgDeal.IndepGuarantee[i] = w
        foundIndepGuarantees[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.IndepGuarantee {
      if !foundIndepGuarantees[i] {
        ctx.orgDeal.IndepGuarantee = append(ctx.orgDeal.IndepGuarantee, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundIndepGuarantees) > 0 {
      var newIndepGuarantees []XsIndepGuaranteeType
      for i, v := range ctx.orgDeal.IndepGuarantee {
        if !foundIndepGuarantees[i] {
          newIndepGuarantees = append(newIndepGuarantees, v)
        }
      }
      ctx.orgDeal.IndepGuarantee = newIndepGuarantees
    }
  }
  foundCollateralInsurances := make(map[int]bool)
  for i, v := range ctx.orgDeal.CollateralInsurance {
    for k, w := range this.CollateralInsurance {
      if w.Same(&v) {
        ctx.orgDeal.CollateralInsurance[i] = w
        foundCollateralInsurances[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.CollateralInsurance {
      if !foundCollateralInsurances[i] {
        ctx.orgDeal.CollateralInsurance = append(ctx.orgDeal.CollateralInsurance, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundCollateralInsurances) > 0 {
      var newCollateralInsurances []XsCollateralInsuranceType
      for i, v := range ctx.orgDeal.CollateralInsurance {
        if !foundCollateralInsurances[i] {
          newCollateralInsurances = append(newCollateralInsurances, v)
        }
      }
      ctx.orgDeal.CollateralInsurance = newCollateralInsurances
    }
  }
  foundSettlements := make(map[int]bool)
  for i, v := range ctx.orgDeal.Settlement {
    for k, w := range this.Settlement {
      if w.Same(&v) {
        ctx.orgDeal.Settlement[i] = w
        foundSettlements[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Settlement {
      if !foundSettlements[i] {
        ctx.orgDeal.Settlement = append(ctx.orgDeal.Settlement, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundSettlements) > 0 {
      var newSettlements []XsSettlementType
      for i, v := range ctx.orgDeal.Settlement {
        if !foundSettlements[i] {
          newSettlements = append(newSettlements, v)
        }
      }
      ctx.orgDeal.Settlement = newSettlements
    }
  }
  if  ctx.op == "D" && this.Repayment.Same(ctx.orgDeal.Repayment)  {
    ctx.orgDeal.Repayment = nil
  }
  if  ctx.op != "D" &&  !this.Repayment.Same(ctx.orgDeal.Repayment) {
    ctx.orgDeal.Repayment = this.Repayment
  }
  return nil
}

func (this *XsOrgDealSecuringChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
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
  return nil
}

// 2.5 Обязательство субъекта прекратилось для денежного обязательства
type XsOrgMonetaryDealEndedTypeFunc func(this *XsOrgMonetaryDealEndedType, ctx *ParseContext) error

var XsOrgMonetaryDealEndedTypeFuncInit XsOrgMonetaryDealEndedTypeFunc = (*XsOrgMonetaryDealEndedType).Init

var XsOrgMonetaryDealEndedTypeFuncLoad XsOrgMonetaryDealEndedTypeFunc = (*XsOrgMonetaryDealEndedType).Load

var XsOrgMonetaryDealEndedTypeFuncValidate XsOrgMonetaryDealEndedTypeFunc = (*XsOrgMonetaryDealEndedType).Validate

func (this *XsOrgMonetaryDealEndedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsTerminationTypeFuncInit(&this.Termination, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncInit(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncInit(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncInit(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncInit(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncInit(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsOrgMonetaryDealEndedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  if   !this.Termination.Same(ctx.orgDeal.Termination) {
    ctx.orgDeal.Termination = &this.Termination
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  if   !this.Amount.Same(ctx.orgDeal.Amount) {
    ctx.orgDeal.Amount = &this.Amount
  }
  if   !this.PaymentTerms.Same(ctx.orgDeal.PaymentTerms) {
    ctx.orgDeal.PaymentTerms = &this.PaymentTerms
  }
  if   !this.FundDate.Same(ctx.orgDeal.FundDate) {
    ctx.orgDeal.FundDate = &this.FundDate
  }
  foundArrear := -1
  for i, v := range ctx.orgDeal.Arrear {
    if this.Arrear.Same(&v) {
      ctx.orgDeal.Arrear[i] = this.Arrear
      foundArrear = i
    }
  }
  if ctx.op != "D" && foundArrear == -1 {
    ctx.orgDeal.Arrear = append(ctx.orgDeal.Arrear, this.Arrear)
  }
  foundDueArrear := -1
  for i, v := range ctx.orgDeal.DueArrear {
    if this.DueArrear.Same(&v) {
      ctx.orgDeal.DueArrear[i] = this.DueArrear
      foundDueArrear = i
    }
  }
  if ctx.op != "D" && foundDueArrear == -1 {
    ctx.orgDeal.DueArrear = append(ctx.orgDeal.DueArrear, this.DueArrear)
  }
  foundPastdueArrear := -1
  for i, v := range ctx.orgDeal.PastdueArrear {
    if this.PastdueArrear.Same(&v) {
      ctx.orgDeal.PastdueArrear[i] = this.PastdueArrear
      foundPastdueArrear = i
    }
  }
  if ctx.op != "D" && foundPastdueArrear == -1 {
    ctx.orgDeal.PastdueArrear = append(ctx.orgDeal.PastdueArrear, this.PastdueArrear)
  }
  foundPaymentss := make(map[int]bool)
  for i, v := range ctx.orgDeal.Payments {
    for k, w := range this.Payments {
      if w.Same(&v) {
        ctx.orgDeal.Payments[i] = w
        foundPaymentss[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Payments {
      if !foundPaymentss[i] {
        ctx.orgDeal.Payments = append(ctx.orgDeal.Payments, w)
      }
    }
  }
  return nil
}

func (this *XsOrgMonetaryDealEndedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsTerminationTypeFuncValidate(&this.Termination, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsPaymentTermsTypeFuncValidate(&this.PaymentTerms, ctx); err != nil { return err }
  if err := XsFundDateTypeFuncValidate(&this.FundDate, ctx); err != nil { return err }
  if err := XsArrearTypeFuncValidate(&this.Arrear, ctx); err != nil { return err }
  if err := XsDueArrearTypeFuncValidate(&this.DueArrear, ctx); err != nil { return err }
  if err := XsPastdueArrearTypeFuncValidate(&this.PastdueArrear, ctx); err != nil { return err }
  for _, v := range (*this).Payments {
    if err := XsPaymentsTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// 2.5 Обязательство субъекта прекратилось для неденежного обязательства
type XsOrgNonMonetaryDealEndedTypeFunc func(this *XsOrgNonMonetaryDealEndedType, ctx *ParseContext) error

var XsOrgNonMonetaryDealEndedTypeFuncInit XsOrgNonMonetaryDealEndedTypeFunc = (*XsOrgNonMonetaryDealEndedType).Init

var XsOrgNonMonetaryDealEndedTypeFuncLoad XsOrgNonMonetaryDealEndedTypeFunc = (*XsOrgNonMonetaryDealEndedType).Load

var XsOrgNonMonetaryDealEndedTypeFuncValidate XsOrgNonMonetaryDealEndedTypeFunc = (*XsOrgNonMonetaryDealEndedType).Validate

func (this *XsOrgNonMonetaryDealEndedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsTerminationTypeFuncInit(&this.Termination, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  if err := XsSubjectNonMonetaryObligationTypeFuncInit(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgNonMonetaryDealEndedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  if   !this.Termination.Same(ctx.orgDeal.Termination) {
    ctx.orgDeal.Termination = &this.Termination
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  if   !this.SubjectNonMonetaryObligation.Same(ctx.orgDeal.SubjectNonMonetaryObligation) {
    ctx.orgDeal.SubjectNonMonetaryObligation = &this.SubjectNonMonetaryObligation
  }
  return nil
}

func (this *XsOrgNonMonetaryDealEndedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsTerminationTypeFuncValidate(&this.Termination, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  if err := XsSubjectNonMonetaryObligationTypeFuncValidate(&this.SubjectNonMonetaryObligation, ctx); err != nil { return err }
  return nil
}

// 2.6 Изменились сведения о судебном споре или требовании по обязательству
type XsOrgDealClaimChangedTypeFunc func(this *XsOrgDealClaimChangedType, ctx *ParseContext) error

var XsOrgDealClaimChangedTypeFuncInit XsOrgDealClaimChangedTypeFunc = (*XsOrgDealClaimChangedType).Init

var XsOrgDealClaimChangedTypeFuncLoad XsOrgDealClaimChangedTypeFunc = (*XsOrgDealClaimChangedType).Load

var XsOrgDealClaimChangedTypeFuncValidate XsOrgDealClaimChangedTypeFunc = (*XsOrgDealClaimChangedType).Validate

func (this *XsOrgDealClaimChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  for _, v := range (*this).Lawsuit {
    if err := XsLawsuitTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsOrgDealClaimChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  foundLawsuits := make(map[int]bool)
  for i, v := range ctx.orgDeal.Lawsuit {
    for k, w := range this.Lawsuit {
      if w.Same(&v) {
        ctx.orgDeal.Lawsuit[i] = w
        foundLawsuits[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Lawsuit {
      if !foundLawsuits[i] {
        ctx.orgDeal.Lawsuit = append(ctx.orgDeal.Lawsuit, w)
      }
    }
  }
  if ctx.op == "D" {
    if len(foundLawsuits) > 0 {
      var newLawsuits []XsLawsuitType
      for i, v := range ctx.orgDeal.Lawsuit {
        if !foundLawsuits[i] {
          newLawsuits = append(newLawsuits, v)
        }
      }
      ctx.orgDeal.Lawsuit = newLawsuits
    }
  }
  return nil
}

func (this *XsOrgDealClaimChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  for _, v := range (*this).Lawsuit {
    if err := XsLawsuitTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// 2.8.1 Конкурсное производство в отношении источника открылось
type XsOrgSourceBankruptcyStartedTypeFunc func(this *XsOrgSourceBankruptcyStartedType, ctx *ParseContext) error

var XsOrgSourceBankruptcyStartedTypeFuncInit XsOrgSourceBankruptcyStartedTypeFunc = (*XsOrgSourceBankruptcyStartedType).Init

var XsOrgSourceBankruptcyStartedTypeFuncLoad XsOrgSourceBankruptcyStartedTypeFunc = (*XsOrgSourceBankruptcyStartedType).Load

var XsOrgSourceBankruptcyStartedTypeFuncValidate XsOrgSourceBankruptcyStartedTypeFunc = (*XsOrgSourceBankruptcyStartedType).Validate

func (this *XsOrgSourceBankruptcyStartedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsSourceBankruptcyTypeFuncInit(&this.SourceBankruptcy, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgSourceBankruptcyStartedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  if   !this.Amount.Same(ctx.orgDeal.Amount) {
    ctx.orgDeal.Amount = &this.Amount
  }
  if  ctx.op == "D" && this.SourceBankruptcy.Same(ctx.orgDeal.SourceBankruptcy)  {
    ctx.orgDeal.SourceBankruptcy = nil
  }
  if  ctx.op != "D" &&  !this.SourceBankruptcy.Same(ctx.orgDeal.SourceBankruptcy) {
    ctx.orgDeal.SourceBankruptcy = &this.SourceBankruptcy
  }
  return nil
}

func (this *XsOrgSourceBankruptcyStartedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsSourceBankruptcyTypeFuncValidate(&this.SourceBankruptcy, ctx); err != nil { return err }
  return nil
}

// 2.8.2 В ходе конкурсного производства изменились сведения об исполнении субъектом своего обязательства или его части либо обязательство прекратилось
type XsOrgSourceBankruptcyChangedTypeFunc func(this *XsOrgSourceBankruptcyChangedType, ctx *ParseContext) error

var XsOrgSourceBankruptcyChangedTypeFuncInit XsOrgSourceBankruptcyChangedTypeFunc = (*XsOrgSourceBankruptcyChangedType).Init

var XsOrgSourceBankruptcyChangedTypeFuncLoad XsOrgSourceBankruptcyChangedTypeFunc = (*XsOrgSourceBankruptcyChangedType).Load

var XsOrgSourceBankruptcyChangedTypeFuncValidate XsOrgSourceBankruptcyChangedTypeFunc = (*XsOrgSourceBankruptcyChangedType).Validate

func (this *XsOrgSourceBankruptcyChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceBankruptcyTypeFuncInit(&this.SourceBankruptcy, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgSourceBankruptcyChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  if   !this.SourceBankruptcy.Same(ctx.orgDeal.SourceBankruptcy) {
    ctx.orgDeal.SourceBankruptcy = &this.SourceBankruptcy
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  return nil
}

func (this *XsOrgSourceBankruptcyChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceBankruptcyTypeFuncValidate(&this.SourceBankruptcy, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 2.8.3 Конкурсное производство в отношении источника завершилось
type XsOrgSourceBankruptcyEndedTypeFunc func(this *XsOrgSourceBankruptcyEndedType, ctx *ParseContext) error

var XsOrgSourceBankruptcyEndedTypeFuncInit XsOrgSourceBankruptcyEndedTypeFunc = (*XsOrgSourceBankruptcyEndedType).Init

var XsOrgSourceBankruptcyEndedTypeFuncLoad XsOrgSourceBankruptcyEndedTypeFunc = (*XsOrgSourceBankruptcyEndedType).Load

var XsOrgSourceBankruptcyEndedTypeFuncValidate XsOrgSourceBankruptcyEndedTypeFunc = (*XsOrgSourceBankruptcyEndedType).Validate

func (this *XsOrgSourceBankruptcyEndedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceBankruptcyTypeFuncInit(&this.SourceBankruptcy, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgSourceBankruptcyEndedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  if   !this.SourceBankruptcy.Same(ctx.orgDeal.SourceBankruptcy) {
    ctx.orgDeal.SourceBankruptcy = &this.SourceBankruptcy
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  return nil
}

func (this *XsOrgSourceBankruptcyEndedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceBankruptcyTypeFuncValidate(&this.SourceBankruptcy, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 2.9.1 Процесс ликвидации источника начался
type XsOrgSourceLiquidationStartedTypeFunc func(this *XsOrgSourceLiquidationStartedType, ctx *ParseContext) error

var XsOrgSourceLiquidationStartedTypeFuncInit XsOrgSourceLiquidationStartedTypeFunc = (*XsOrgSourceLiquidationStartedType).Init

var XsOrgSourceLiquidationStartedTypeFuncLoad XsOrgSourceLiquidationStartedTypeFunc = (*XsOrgSourceLiquidationStartedType).Load

var XsOrgSourceLiquidationStartedTypeFuncValidate XsOrgSourceLiquidationStartedTypeFunc = (*XsOrgSourceLiquidationStartedType).Validate

func (this *XsOrgSourceLiquidationStartedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncInit(&this.Amount, ctx); err != nil { return err }
  if err := XsSourceLiquidationTypeFuncInit(&this.SourceLiquidation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgSourceLiquidationStartedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  err := saveSubjectRole(this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  if   !this.Amount.Same(ctx.orgDeal.Amount) {
    ctx.orgDeal.Amount = &this.Amount
  }
  if  ctx.op == "D" && this.SourceLiquidation.Same(ctx.orgDeal.SourceLiquidation)  {
    ctx.orgDeal.SourceLiquidation = nil
  }
  if  ctx.op != "D" &&  !this.SourceLiquidation.Same(ctx.orgDeal.SourceLiquidation) {
    ctx.orgDeal.SourceLiquidation = &this.SourceLiquidation
  }
  return nil
}

func (this *XsOrgSourceLiquidationStartedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if err := XsDealAmountTypeFuncValidate(&this.Amount, ctx); err != nil { return err }
  if err := XsSourceLiquidationTypeFuncValidate(&this.SourceLiquidation, ctx); err != nil { return err }
  return nil
}

// 2.9.2 В ходе процесса ликвидации изменились сведения об исполнении субъектом своего обязательства или его части либо обязательство прекратилось
type XsOrgSourceLiquidationChangedTypeFunc func(this *XsOrgSourceLiquidationChangedType, ctx *ParseContext) error

var XsOrgSourceLiquidationChangedTypeFuncInit XsOrgSourceLiquidationChangedTypeFunc = (*XsOrgSourceLiquidationChangedType).Init

var XsOrgSourceLiquidationChangedTypeFuncLoad XsOrgSourceLiquidationChangedTypeFunc = (*XsOrgSourceLiquidationChangedType).Load

var XsOrgSourceLiquidationChangedTypeFuncValidate XsOrgSourceLiquidationChangedTypeFunc = (*XsOrgSourceLiquidationChangedType).Validate

func (this *XsOrgSourceLiquidationChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceLiquidationTypeFuncInit(&this.SourceLiquidation, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgSourceLiquidationChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  if   !this.SourceLiquidation.Same(ctx.orgDeal.SourceLiquidation) {
    ctx.orgDeal.SourceLiquidation = &this.SourceLiquidation
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  return nil
}

func (this *XsOrgSourceLiquidationChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceLiquidationTypeFuncValidate(&this.SourceLiquidation, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 2.9.3 Процесс ликвидации источника завершился
type XsOrgSourceLiquidationEndedTypeFunc func(this *XsOrgSourceLiquidationEndedType, ctx *ParseContext) error

var XsOrgSourceLiquidationEndedTypeFuncInit XsOrgSourceLiquidationEndedTypeFunc = (*XsOrgSourceLiquidationEndedType).Init

var XsOrgSourceLiquidationEndedTypeFuncLoad XsOrgSourceLiquidationEndedTypeFunc = (*XsOrgSourceLiquidationEndedType).Load

var XsOrgSourceLiquidationEndedTypeFuncValidate XsOrgSourceLiquidationEndedTypeFunc = (*XsOrgSourceLiquidationEndedType).Validate

func (this *XsOrgSourceLiquidationEndedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceLiquidationTypeFuncInit(&this.SourceLiquidation, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncInit(&this.Participation, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgSourceLiquidationEndedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  if   !this.SourceLiquidation.Same(ctx.orgDeal.SourceLiquidation) {
    ctx.orgDeal.SourceLiquidation = &this.SourceLiquidation
  }
  foundParticipation := -1
  for i, v := range ctx.org.Participation {
    if this.Participation.Same(&v) {
      ctx.org.Participation[i] = this.Participation
      foundParticipation = i
    }
  }
  if ctx.op != "D" && foundParticipation == -1 {
    ctx.org.Participation = append(ctx.org.Participation, this.Participation)
  }
  return nil
}

func (this *XsOrgSourceLiquidationEndedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsSourceLiquidationTypeFuncValidate(&this.SourceLiquidation, ctx); err != nil { return err }
  if err := XsParticipationTypeFuncValidate(&this.Participation, ctx); err != nil { return err }
  return nil
}

// 2.10 Источник прекратил передачу информации по обязательству
type XsOrgDealInfoStoppedTypeFunc func(this *XsOrgDealInfoStoppedType, ctx *ParseContext) error

var XsOrgDealInfoStoppedTypeFuncInit XsOrgDealInfoStoppedTypeFunc = (*XsOrgDealInfoStoppedType).Init

var XsOrgDealInfoStoppedTypeFuncLoad XsOrgDealInfoStoppedTypeFunc = (*XsOrgDealInfoStoppedType).Load

var XsOrgDealInfoStoppedTypeFuncValidate XsOrgDealInfoStoppedTypeFunc = (*XsOrgDealInfoStoppedType).Validate

func (this *XsOrgDealInfoStoppedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsStopInfoTypeFuncInit(&this.StopInfo, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgDealInfoStoppedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  foundStopInfo := -1
  for i, v := range ctx.orgDeal.StopInfo {
    if this.StopInfo.Same(&v) {
      ctx.orgDeal.StopInfo[i] = this.StopInfo
      foundStopInfo = i
    }
  }
  if ctx.op != "D" && foundStopInfo == -1 {
    ctx.orgDeal.StopInfo = append(ctx.orgDeal.StopInfo, this.StopInfo)
  }
  if ctx.op == "D" && foundStopInfo != -1 {
    ctx.orgDeal.StopInfo = append(ctx.orgDeal.StopInfo[:foundStopInfo], ctx.orgDeal.StopInfo[foundStopInfo+1:]...)
  }
  return nil
}

func (this *XsOrgDealInfoStoppedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsStopInfoTypeFuncValidate(&this.StopInfo, ctx); err != nil { return err }
  return nil
}

// 2.11 Права кредитора по обязательству перешли к другому лицу
type XsOrgCreditorChangedTypeFunc func(this *XsOrgCreditorChangedType, ctx *ParseContext) error

var XsOrgCreditorChangedTypeFuncInit XsOrgCreditorChangedTypeFunc = (*XsOrgCreditorChangedType).Init

var XsOrgCreditorChangedTypeFuncLoad XsOrgCreditorChangedTypeFunc = (*XsOrgCreditorChangedType).Load

var XsOrgCreditorChangedTypeFuncValidate XsOrgCreditorChangedTypeFunc = (*XsOrgCreditorChangedType).Validate

func (this *XsOrgCreditorChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsStopInfoTypeFuncInit(&this.StopInfo, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncInit(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncInit(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsOrgCreditorChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  foundStopInfo := -1
  for i, v := range ctx.orgDeal.StopInfo {
    if this.StopInfo.Same(&v) {
      ctx.orgDeal.StopInfo[i] = this.StopInfo
      foundStopInfo = i
    }
  }
  if ctx.op != "D" && foundStopInfo == -1 {
    ctx.orgDeal.StopInfo = append(ctx.orgDeal.StopInfo, this.StopInfo)
  }
  if   !this.Accounting.Same(ctx.orgDeal.Accounting) {
    ctx.orgDeal.Accounting = &this.Accounting
  }
  if  ctx.op == "D" && this.OrgAcquirerInfo.Same(ctx.orgDeal.OrgAcquirerInfo)  {
    ctx.orgDeal.OrgAcquirerInfo = nil
  }
  if  ctx.op != "D" &&  !this.OrgAcquirerInfo.Same(ctx.orgDeal.OrgAcquirerInfo) {
    ctx.orgDeal.OrgAcquirerInfo = this.OrgAcquirerInfo
  }
  if  ctx.op == "D" && this.PersonAcquirerInfo.Same(ctx.orgDeal.PersonAcquirerInfo)  {
    ctx.orgDeal.PersonAcquirerInfo = nil
  }
  if  ctx.op != "D" &&  !this.PersonAcquirerInfo.Same(ctx.orgDeal.PersonAcquirerInfo) {
    ctx.orgDeal.PersonAcquirerInfo = this.PersonAcquirerInfo
  }
  return nil
}

func (this *XsOrgCreditorChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsStopInfoTypeFuncValidate(&this.StopInfo, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncValidate(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncValidate(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  return nil
}

// 2.12 Изменились сведения об обслуживающей организации (в частности, заключен, изменен или расторгнут договор обслуживания)
type XsOrgServiceOrgChangedTypeFunc func(this *XsOrgServiceOrgChangedType, ctx *ParseContext) error

var XsOrgServiceOrgChangedTypeFuncInit XsOrgServiceOrgChangedTypeFunc = (*XsOrgServiceOrgChangedType).Init

var XsOrgServiceOrgChangedTypeFuncLoad XsOrgServiceOrgChangedTypeFunc = (*XsOrgServiceOrgChangedType).Load

var XsOrgServiceOrgChangedTypeFuncValidate XsOrgServiceOrgChangedTypeFunc = (*XsOrgServiceOrgChangedType).Validate

func (this *XsOrgServiceOrgChangedType) Init(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsServiceOrgTypeFuncInit(&this.ServiceOrg, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncInit(&this.Accounting, ctx); err != nil { return err }
  return nil
}

func (this *XsOrgServiceOrgChangedType) Load(ctx *ParseContext) error {
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  if  ctx.op == "D" && this.ServiceOrg.Same(ctx.orgDeal.ServiceOrg)  {
    ctx.orgDeal.ServiceOrg = nil
  }
  if  ctx.op != "D" &&  !this.ServiceOrg.Same(ctx.orgDeal.ServiceOrg) {
    ctx.orgDeal.ServiceOrg = &this.ServiceOrg
  }
  if   !this.Accounting.Same(ctx.orgDeal.Accounting) {
    ctx.orgDeal.Accounting = &this.Accounting
  }
  return nil
}

func (this *XsOrgServiceOrgChangedType) Validate(ctx *ParseContext) error {
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsServiceOrgTypeFuncValidate(&this.ServiceOrg, ctx); err != nil { return err }
  if err := XsAccountingTypeFuncValidate(&this.Accounting, ctx); err != nil { return err }
  return nil
}

// Внутреннее событие Кредитной - используется только Кредитным Бюро
type XsInternalEventTypeFunc func(this *XsInternalEventType, ctx *ParseContext) error

var XsInternalEventTypeFuncInit XsInternalEventTypeFunc = (*XsInternalEventType).Init

var XsInternalEventTypeFuncLoad XsInternalEventTypeFunc = (*XsInternalEventType).Load

var XsInternalEventTypeFuncValidate XsInternalEventTypeFunc = (*XsInternalEventType).Validate

func (this *XsInternalEventType) Init(ctx *ParseContext) error {
  if this.Subject != nil {
    if err := XsPersonSubjectTypeFuncInit(this.Subject, ctx); err != nil { return err }
  }
  if this.RatingCalculated != nil {
    if err := XsRatingCalculatedTypeFuncInit(this.RatingCalculated, ctx); err != nil { return err }
  }
  if this.ScoringRequested != nil {
    if err := XsScoringRequestedTypeFuncInit(this.ScoringRequested, ctx); err != nil { return err }
  }
  if this.ReportRequested != nil {
    if err := XsReportRequestedTypeFuncInit(this.ReportRequested, ctx); err != nil { return err }
  }
  if this.AvgPaymentsSupplied != nil {
    if err := XsAvgPaymentsSuppliedTypeFuncInit(this.AvgPaymentsSupplied, ctx); err != nil { return err }
  }
  if this.ImportPerson != nil {
    if err := XsImportPersonTypeFuncInit(this.ImportPerson, ctx); err != nil { return err }
  }
  if this.ImportOrg != nil {
    if err := XsImportOrgTypeFuncInit(this.ImportOrg, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsInternalEventType) Load(ctx *ParseContext) error {
  if this.Subject != nil {
    if err := XsPersonSubjectTypeFuncLoad(this.Subject, ctx); err != nil { return err }
  }
  if this.RatingCalculated != nil {
    if err := XsRatingCalculatedTypeFuncLoad(this.RatingCalculated, ctx); err != nil { return err }
  }
  if this.ScoringRequested != nil {
    if err := XsScoringRequestedTypeFuncLoad(this.ScoringRequested, ctx); err != nil { return err }
  }
  if this.ReportRequested != nil {
    if err := XsReportRequestedTypeFuncLoad(this.ReportRequested, ctx); err != nil { return err }
  }
  if this.AvgPaymentsSupplied != nil {
    if err := XsAvgPaymentsSuppliedTypeFuncLoad(this.AvgPaymentsSupplied, ctx); err != nil { return err }
  }
  if this.ImportPerson != nil {
    if err := XsImportPersonTypeFuncLoad(this.ImportPerson, ctx); err != nil { return err }
  }
  if this.ImportOrg != nil {
    if err := XsImportOrgTypeFuncLoad(this.ImportOrg, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsInternalEventType) Validate(ctx *ParseContext) error {
  if this.Subject != nil {
    if err := XsPersonSubjectTypeFuncValidate(this.Subject, ctx); err != nil { return err }
  }
  if this.RatingCalculated != nil {
    if err := XsRatingCalculatedTypeFuncValidate(this.RatingCalculated, ctx); err != nil { return err }
  }
  if this.ScoringRequested != nil {
    if err := XsScoringRequestedTypeFuncValidate(this.ScoringRequested, ctx); err != nil { return err }
  }
  if this.ReportRequested != nil {
    if err := XsReportRequestedTypeFuncValidate(this.ReportRequested, ctx); err != nil { return err }
  }
  if this.AvgPaymentsSupplied != nil {
    if err := XsAvgPaymentsSuppliedTypeFuncValidate(this.AvgPaymentsSupplied, ctx); err != nil { return err }
  }
  if this.ImportPerson != nil {
    if err := XsImportPersonTypeFuncValidate(this.ImportPerson, ctx); err != nil { return err }
  }
  if this.ImportOrg != nil {
    if err := XsImportOrgTypeFuncValidate(this.ImportOrg, ctx); err != nil { return err }
  }
  return nil
}

// Импорт КИ физического лица
type XsImportPersonTypeFunc func(this *XsImportPersonType, ctx *ParseContext) error

var XsImportPersonTypeFuncInit XsImportPersonTypeFunc = (*XsImportPersonType).Init

var XsImportPersonTypeFuncLoad XsImportPersonTypeFunc = (*XsImportPersonType).Load

var XsImportPersonTypeFuncValidate XsImportPersonTypeFunc = (*XsImportPersonType).Validate

func (this *XsImportPersonType) Init(ctx *ParseContext) error {
  if err := XsPersonNameTypeFuncInit(&this.Name, ctx); err != nil { return err }
  for _, v := range (*this).PrevName {
    if err := XsPersonPrevNameTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsBirthInfoTypeFuncInit(&this.BirthInfo, ctx); err != nil { return err }
  for _, v := range (*this).Id {
    if err := XsPersonIdTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevId {
    if err := XsPersonPrevIdTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Tax {
    if err := XsTaxTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.SocialId != nil {
    if err := XsSocialIdTypeFuncInit(this.SocialId, ctx); err != nil { return err }
  }
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
  for _, v := range (*this).Rating {
    if err := XsRatingTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Score {
    if err := XsScoreTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
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
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncInit(this.FundDate, ctx); err != nil { return err }
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
  for _, v := range (*this).AvgPayment {
    if err := XsAvgPaymentTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.SourceBankruptcy != nil {
    if err := XsSourceBankruptcyTypeFuncInit(this.SourceBankruptcy, ctx); err != nil { return err }
  }
  if this.SourceLiquidation != nil {
    if err := XsSourceLiquidationTypeFuncInit(this.SourceLiquidation, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collection {
    if err := XsCollectionTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).CustomerRequest {
    if err := XsCustomerRequestTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).StopInfo {
    if err := XsStopInfoTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncInit(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncInit(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncInit(this.CommissionerSource, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgCustomer {
    if err := XsOrgCustomerTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PersonCustomer {
    if err := XsPersonCustomerTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncInit(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncInit(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  if this.ServiceOrg != nil {
    if err := XsServiceOrgTypeFuncInit(this.ServiceOrg, ctx); err != nil { return err }
  }
  if this.Accounting != nil {
    if err := XsAccountingTypeFuncInit(this.Accounting, ctx); err != nil { return err }
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
  return nil
}

func (this *XsImportPersonType) Load(ctx *ParseContext) error {
  if   !this.Name.Same(&ctx.person.Name) {
    ctx.person.Name = this.Name
  }
  foundPrevNames := make(map[int]bool)
  for i, v := range ctx.person.PrevName {
    for k, w := range this.PrevName {
      if w.Same(&v) {
        ctx.person.PrevName[i] = w
        foundPrevNames[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevName {
      if !foundPrevNames[i] {
        ctx.person.PrevName = append(ctx.person.PrevName, w)
      }
    }
  }
  if   !this.BirthInfo.Same(&ctx.person.BirthInfo) {
    ctx.person.BirthInfo = this.BirthInfo
  }
  foundIds := make(map[int]bool)
  for i, v := range ctx.person.Id {
    for k, w := range this.Id {
      if w.Same(&v) {
        ctx.person.Id[i] = w
        foundIds[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Id {
      if !foundIds[i] {
        ctx.person.Id = append(ctx.person.Id, w)
      }
    }
  }
  foundPrevIds := make(map[int]bool)
  for i, v := range ctx.person.PrevId {
    for k, w := range this.PrevId {
      if w.Same(&v) {
        ctx.person.PrevId[i] = w
        foundPrevIds[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevId {
      if !foundPrevIds[i] {
        ctx.person.PrevId = append(ctx.person.PrevId, w)
      }
    }
  }
  foundTaxs := make(map[int]bool)
  for i, v := range ctx.person.Tax {
    for k, w := range this.Tax {
      if w.Same(&v) {
        ctx.person.Tax[i] = w
        foundTaxs[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Tax {
      if !foundTaxs[i] {
        ctx.person.Tax = append(ctx.person.Tax, w)
      }
    }
  }
  if   !this.SocialId.Same(ctx.person.SocialId) {
    ctx.person.SocialId = this.SocialId
  }
  if   !this.Address.Same(&ctx.person.Address) {
    ctx.person.Address = this.Address
  }
  if   !this.FactAddress.Same(ctx.person.FactAddress) {
    ctx.person.FactAddress = this.FactAddress
  }
  foundContacts := make(map[int]bool)
  for i, v := range ctx.person.Contact {
    for k, w := range this.Contact {
      if w.Same(&v) {
        ctx.person.Contact[i] = w
        foundContacts[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Contact {
      if !foundContacts[i] {
        ctx.person.Contact = append(ctx.person.Contact, w)
      }
    }
  }
  if   !this.SoleProprietor.Same(ctx.person.SoleProprietor) {
    ctx.person.SoleProprietor = this.SoleProprietor
  }
  foundLegalCapacitys := make(map[int]bool)
  for i, v := range ctx.person.LegalCapacity {
    for k, w := range this.LegalCapacity {
      if w.Same(&v) {
        ctx.person.LegalCapacity[i] = w
        foundLegalCapacitys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.LegalCapacity {
      if !foundLegalCapacitys[i] {
        ctx.person.LegalCapacity = append(ctx.person.LegalCapacity, w)
      }
    }
  }
  foundBankruptcys := make(map[int]bool)
  for i, v := range ctx.person.Bankruptcy {
    for k, w := range this.Bankruptcy {
      if w.Same(&v) {
        ctx.person.Bankruptcy[i] = w
        foundBankruptcys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Bankruptcy {
      if !foundBankruptcys[i] {
        ctx.person.Bankruptcy = append(ctx.person.Bankruptcy, w)
      }
    }
  }
  foundFulfillments := make(map[int]bool)
  for i, v := range ctx.person.Fulfillment {
    for k, w := range this.Fulfillment {
      if w.Same(&v) {
        ctx.person.Fulfillment[i] = w
        foundFulfillments[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Fulfillment {
      if !foundFulfillments[i] {
        ctx.person.Fulfillment = append(ctx.person.Fulfillment, w)
      }
    }
  }
  foundRatings := make(map[int]bool)
  for i, v := range ctx.person.Rating {
    for k, w := range this.Rating {
      if w.Same(&v) {
        ctx.person.Rating[i] = w
        foundRatings[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Rating {
      if !foundRatings[i] {
        ctx.person.Rating = append(ctx.person.Rating, w)
      }
    }
  }
  foundScores := make(map[int]bool)
  for i, v := range ctx.person.Score {
    for k, w := range this.Score {
      if w.Same(&v) {
        ctx.person.Score[i] = w
        foundScores[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Score {
      if !foundScores[i] {
        ctx.person.Score = append(ctx.person.Score, w)
      }
    }
  }
  if   !this.DealUid.Same(&ctx.personDeal.DealUid) {
    ctx.personDeal.DealUid = this.DealUid
  }
  if   !this.Main.Same(&ctx.personDeal.Main) {
    ctx.personDeal.Main = this.Main
  }
  if   !this.Amount.Same(ctx.personDeal.Amount) {
    ctx.personDeal.Amount = this.Amount
  }
  if   !this.JointDebtors.Same(ctx.personDeal.JointDebtors) {
    ctx.personDeal.JointDebtors = this.JointDebtors
  }
  if   !this.PaymentTerms.Same(ctx.personDeal.PaymentTerms) {
    ctx.personDeal.PaymentTerms = this.PaymentTerms
  }
  if   !this.TotalCost.Same(ctx.personDeal.TotalCost) {
    ctx.personDeal.TotalCost = this.TotalCost
  }
  foundChanges := make(map[int]bool)
  for i, v := range ctx.personDeal.Change {
    for k, w := range this.Change {
      if w.Same(&v) {
        ctx.personDeal.Change[i] = w
        foundChanges[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Change {
      if !foundChanges[i] {
        ctx.personDeal.Change = append(ctx.personDeal.Change, w)
      }
    }
  }
  if   !this.FundDate.Same(ctx.personDeal.FundDate) {
    ctx.personDeal.FundDate = this.FundDate
  }
  foundArrears := make(map[int]bool)
  for i, v := range ctx.personDeal.Arrear {
    for k, w := range this.Arrear {
      if w.Same(&v) {
        ctx.personDeal.Arrear[i] = w
        foundArrears[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Arrear {
      if !foundArrears[i] {
        ctx.personDeal.Arrear = append(ctx.personDeal.Arrear, w)
      }
    }
  }
  foundDueArrears := make(map[int]bool)
  for i, v := range ctx.personDeal.DueArrear {
    for k, w := range this.DueArrear {
      if w.Same(&v) {
        ctx.personDeal.DueArrear[i] = w
        foundDueArrears[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.DueArrear {
      if !foundDueArrears[i] {
        ctx.personDeal.DueArrear = append(ctx.personDeal.DueArrear, w)
      }
    }
  }
  foundPastdueArrears := make(map[int]bool)
  for i, v := range ctx.personDeal.PastdueArrear {
    for k, w := range this.PastdueArrear {
      if w.Same(&v) {
        ctx.personDeal.PastdueArrear[i] = w
        foundPastdueArrears[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PastdueArrear {
      if !foundPastdueArrears[i] {
        ctx.personDeal.PastdueArrear = append(ctx.personDeal.PastdueArrear, w)
      }
    }
  }
  foundPaymentss := make(map[int]bool)
  for i, v := range ctx.personDeal.Payments {
    for k, w := range this.Payments {
      if w.Same(&v) {
        ctx.personDeal.Payments[i] = w
        foundPaymentss[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Payments {
      if !foundPaymentss[i] {
        ctx.personDeal.Payments = append(ctx.personDeal.Payments, w)
      }
    }
  }
  if   !this.MonthlyPayment.Same(ctx.personDeal.MonthlyPayment) {
    ctx.personDeal.MonthlyPayment = this.MonthlyPayment
  }
  if   !this.SourceNonMonetaryObligation.Same(ctx.personDeal.SourceNonMonetaryObligation) {
    ctx.personDeal.SourceNonMonetaryObligation = this.SourceNonMonetaryObligation
  }
  if   !this.SubjectNonMonetaryObligation.Same(ctx.personDeal.SubjectNonMonetaryObligation) {
    ctx.personDeal.SubjectNonMonetaryObligation = this.SubjectNonMonetaryObligation
  }
  foundCollaterals := make(map[int]bool)
  for i, v := range ctx.personDeal.Collateral {
    for k, w := range this.Collateral {
      if w.Same(&v) {
        ctx.personDeal.Collateral[i] = w
        foundCollaterals[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Collateral {
      if !foundCollaterals[i] {
        ctx.personDeal.Collateral = append(ctx.personDeal.Collateral, w)
      }
    }
  }
  foundWarrantys := make(map[int]bool)
  for i, v := range ctx.personDeal.Warranty {
    for k, w := range this.Warranty {
      if w.Same(&v) {
        ctx.personDeal.Warranty[i] = w
        foundWarrantys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Warranty {
      if !foundWarrantys[i] {
        ctx.personDeal.Warranty = append(ctx.personDeal.Warranty, w)
      }
    }
  }
  foundIndepGuarantees := make(map[int]bool)
  for i, v := range ctx.personDeal.IndepGuarantee {
    for k, w := range this.IndepGuarantee {
      if w.Same(&v) {
        ctx.personDeal.IndepGuarantee[i] = w
        foundIndepGuarantees[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.IndepGuarantee {
      if !foundIndepGuarantees[i] {
        ctx.personDeal.IndepGuarantee = append(ctx.personDeal.IndepGuarantee, w)
      }
    }
  }
  foundCollateralInsurances := make(map[int]bool)
  for i, v := range ctx.personDeal.CollateralInsurance {
    for k, w := range this.CollateralInsurance {
      if w.Same(&v) {
        ctx.personDeal.CollateralInsurance[i] = w
        foundCollateralInsurances[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.CollateralInsurance {
      if !foundCollateralInsurances[i] {
        ctx.personDeal.CollateralInsurance = append(ctx.personDeal.CollateralInsurance, w)
      }
    }
  }
  foundSettlements := make(map[int]bool)
  for i, v := range ctx.personDeal.Settlement {
    for k, w := range this.Settlement {
      if w.Same(&v) {
        ctx.personDeal.Settlement[i] = w
        foundSettlements[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Settlement {
      if !foundSettlements[i] {
        ctx.personDeal.Settlement = append(ctx.personDeal.Settlement, w)
      }
    }
  }
  if   !this.Repayment.Same(ctx.personDeal.Repayment) {
    ctx.personDeal.Repayment = this.Repayment
  }
  if   !this.Termination.Same(ctx.personDeal.Termination) {
    ctx.personDeal.Termination = this.Termination
  }
  foundLawsuits := make(map[int]bool)
  for i, v := range ctx.personDeal.Lawsuit {
    for k, w := range this.Lawsuit {
      if w.Same(&v) {
        ctx.personDeal.Lawsuit[i] = w
        foundLawsuits[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Lawsuit {
      if !foundLawsuits[i] {
        ctx.personDeal.Lawsuit = append(ctx.personDeal.Lawsuit, w)
      }
    }
  }
  foundAvgPayments := make(map[int]bool)
  for i, v := range ctx.person.AvgPayment {
    for k, w := range this.AvgPayment {
      if w.Same(&v) {
        ctx.person.AvgPayment[i] = w
        foundAvgPayments[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.AvgPayment {
      if !foundAvgPayments[i] {
        ctx.person.AvgPayment = append(ctx.person.AvgPayment, w)
      }
    }
  }
  if   !this.SourceBankruptcy.Same(ctx.personDeal.SourceBankruptcy) {
    ctx.personDeal.SourceBankruptcy = this.SourceBankruptcy
  }
  if   !this.SourceLiquidation.Same(ctx.personDeal.SourceLiquidation) {
    ctx.personDeal.SourceLiquidation = this.SourceLiquidation
  }
  foundCollections := make(map[int]bool)
  for i, v := range ctx.person.Collection {
    for k, w := range this.Collection {
      if w.Same(&v) {
        ctx.person.Collection[i] = w
        foundCollections[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Collection {
      if !foundCollections[i] {
        ctx.person.Collection = append(ctx.person.Collection, w)
      }
    }
  }
  foundStopInfos := make(map[int]bool)
  for i, v := range ctx.personDeal.StopInfo {
    for k, w := range this.StopInfo {
      if w.Same(&v) {
        ctx.personDeal.StopInfo[i] = w
        foundStopInfos[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.StopInfo {
      if !foundStopInfos[i] {
        ctx.personDeal.StopInfo = append(ctx.personDeal.StopInfo, w)
      }
    }
  }
  foundOrgCustomers := make(map[int]bool)
  for i, v := range ctx.person.OrgCustomer {
    for k, w := range this.OrgCustomer {
      if w.Same(&v) {
        ctx.person.OrgCustomer[i] = w
        foundOrgCustomers[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.OrgCustomer {
      if !foundOrgCustomers[i] {
        ctx.person.OrgCustomer = append(ctx.person.OrgCustomer, w)
      }
    }
  }
  foundPersonCustomers := make(map[int]bool)
  for i, v := range ctx.person.PersonCustomer {
    for k, w := range this.PersonCustomer {
      if w.Same(&v) {
        ctx.person.PersonCustomer[i] = w
        foundPersonCustomers[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PersonCustomer {
      if !foundPersonCustomers[i] {
        ctx.person.PersonCustomer = append(ctx.person.PersonCustomer, w)
      }
    }
  }
  if   !this.OrgAcquirerInfo.Same(ctx.personDeal.OrgAcquirerInfo) {
    ctx.personDeal.OrgAcquirerInfo = this.OrgAcquirerInfo
  }
  if   !this.PersonAcquirerInfo.Same(ctx.personDeal.PersonAcquirerInfo) {
    ctx.personDeal.PersonAcquirerInfo = this.PersonAcquirerInfo
  }
  if   !this.ServiceOrg.Same(ctx.personDeal.ServiceOrg) {
    ctx.personDeal.ServiceOrg = this.ServiceOrg
  }
  if   !this.Accounting.Same(ctx.personDeal.Accounting) {
    ctx.personDeal.Accounting = this.Accounting
  }
  foundApplications := make(map[int]bool)
  for i, v := range ctx.person.Application {
    for k, w := range this.Application {
      if w.Same(&v) {
        ctx.person.Application[i] = w
        foundApplications[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Application {
      if !foundApplications[i] {
        ctx.person.Application = append(ctx.person.Application, w)
      }
    }
  }
  foundParticipations := make(map[int]bool)
  for i, v := range ctx.person.Participation {
    for k, w := range this.Participation {
      if w.Same(&v) {
        ctx.person.Participation[i] = w
        foundParticipations[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Participation {
      if !foundParticipations[i] {
        ctx.person.Participation = append(ctx.person.Participation, w)
      }
    }
  }
  foundAppRejects := make(map[int]bool)
  for i, v := range ctx.person.AppReject {
    for k, w := range this.AppReject {
      if w.Same(&v) {
        ctx.person.AppReject[i] = w
        foundAppRejects[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.AppReject {
      if !foundAppRejects[i] {
        ctx.person.AppReject = append(ctx.person.AppReject, w)
      }
    }
  }
  err := saveSubjectRole(*this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  return nil
}

func (this *XsImportPersonType) Validate(ctx *ParseContext) error {
  if err := XsPersonNameTypeFuncValidate(&this.Name, ctx); err != nil { return err }
  for _, v := range (*this).PrevName {
    if err := XsPersonPrevNameTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsBirthInfoTypeFuncValidate(&this.BirthInfo, ctx); err != nil { return err }
  for _, v := range (*this).Id {
    if err := XsPersonIdTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevId {
    if err := XsPersonPrevIdTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Tax {
    if err := XsTaxTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.SocialId != nil {
    if err := XsSocialIdTypeFuncValidate(this.SocialId, ctx); err != nil { return err }
  }
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
  for _, v := range (*this).Rating {
    if err := XsRatingTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Score {
    if err := XsScoreTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
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
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncValidate(this.FundDate, ctx); err != nil { return err }
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
  for _, v := range (*this).AvgPayment {
    if err := XsAvgPaymentTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.SourceBankruptcy != nil {
    if err := XsSourceBankruptcyTypeFuncValidate(this.SourceBankruptcy, ctx); err != nil { return err }
  }
  if this.SourceLiquidation != nil {
    if err := XsSourceLiquidationTypeFuncValidate(this.SourceLiquidation, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collection {
    if err := XsCollectionTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).CustomerRequest {
    if err := XsCustomerRequestTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).StopInfo {
    if err := XsStopInfoTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncValidate(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncValidate(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncValidate(this.CommissionerSource, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgCustomer {
    if err := XsOrgCustomerTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PersonCustomer {
    if err := XsPersonCustomerTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncValidate(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncValidate(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  if this.ServiceOrg != nil {
    if err := XsServiceOrgTypeFuncValidate(this.ServiceOrg, ctx); err != nil { return err }
  }
  if this.Accounting != nil {
    if err := XsAccountingTypeFuncValidate(this.Accounting, ctx); err != nil { return err }
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
  return nil
}

// Импорт КИ юридического лица
type XsImportOrgTypeFunc func(this *XsImportOrgType, ctx *ParseContext) error

var XsImportOrgTypeFuncInit XsImportOrgTypeFunc = (*XsImportOrgType).Init

var XsImportOrgTypeFuncLoad XsImportOrgTypeFunc = (*XsImportOrgType).Load

var XsImportOrgTypeFuncValidate XsImportOrgTypeFunc = (*XsImportOrgType).Validate

func (this *XsImportOrgType) Init(ctx *ParseContext) error {
  if err := XsOrgNameTypeFuncInit(&this.Name, ctx); err != nil { return err }
  if err := XsOrgAddressTypeFuncInit(&this.Address, ctx); err != nil { return err }
  if err := XsOrgRegNumTypeFuncInit(&this.RegNum, ctx); err != nil { return err }
  if this.Tax != nil {
    if err := XsOrgTaxTypeFuncInit(this.Tax, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgChange {
    if err := XsOrgChangeTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Bankruptcy {
    if err := XsBankruptcyTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Fulfillment {
    if err := XsOrgFulfillmentTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Score {
    if err := XsScoreTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if this.Amount != nil {
    if err := XsDealAmountTypeFuncInit(this.Amount, ctx); err != nil { return err }
  }
  if this.JointDebtors != nil {
    if err := XsJointDebtorsTypeFuncInit(this.JointDebtors, ctx); err != nil { return err }
  }
  if this.PaymentTerms != nil {
    if err := XsPaymentTermsTypeFuncInit(this.PaymentTerms, ctx); err != nil { return err }
  }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncInit(this.FundDate, ctx); err != nil { return err }
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
  if this.Collection != nil {
    if err := XsCollectionTypeFuncInit(this.Collection, ctx); err != nil { return err }
  }
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncInit(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncInit(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncInit(this.CommissionerSource, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgCustomer {
    if err := XsOrgCustomerTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PersonCustomer {
    if err := XsPersonCustomerTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncInit(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncInit(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  if this.ServiceOrg != nil {
    if err := XsServiceOrgTypeFuncInit(this.ServiceOrg, ctx); err != nil { return err }
  }
  if this.Accounting != nil {
    if err := XsAccountingTypeFuncInit(this.Accounting, ctx); err != nil { return err }
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
  return nil
}

func (this *XsImportOrgType) Load(ctx *ParseContext) error {
  if   !this.Name.Same(&ctx.org.Name) {
    ctx.org.Name = this.Name
  }
  if   !this.Address.Same(&ctx.org.Address) {
    ctx.org.Address = this.Address
  }
  if   !this.RegNum.Same(&ctx.org.RegNum) {
    ctx.org.RegNum = this.RegNum
  }
  if   !this.Tax.Same(ctx.org.Tax) {
    ctx.org.Tax = this.Tax
  }
  foundOrgChanges := make(map[int]bool)
  for i, v := range ctx.org.OrgChange {
    for k, w := range this.OrgChange {
      if w.Same(&v) {
        ctx.org.OrgChange[i] = w
        foundOrgChanges[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.OrgChange {
      if !foundOrgChanges[i] {
        ctx.org.OrgChange = append(ctx.org.OrgChange, w)
      }
    }
  }
  foundBankruptcys := make(map[int]bool)
  for i, v := range ctx.org.Bankruptcy {
    for k, w := range this.Bankruptcy {
      if w.Same(&v) {
        ctx.org.Bankruptcy[i] = w
        foundBankruptcys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Bankruptcy {
      if !foundBankruptcys[i] {
        ctx.org.Bankruptcy = append(ctx.org.Bankruptcy, w)
      }
    }
  }
  foundFulfillments := make(map[int]bool)
  for i, v := range ctx.org.Fulfillment {
    for k, w := range this.Fulfillment {
      if w.Same(&v) {
        ctx.org.Fulfillment[i] = w
        foundFulfillments[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Fulfillment {
      if !foundFulfillments[i] {
        ctx.org.Fulfillment = append(ctx.org.Fulfillment, w)
      }
    }
  }
  foundPrevOrgs := make(map[int]bool)
  for i, v := range ctx.org.PrevOrg {
    for k, w := range this.PrevOrg {
      if w.Same(&v) {
        ctx.org.PrevOrg[i] = w
        foundPrevOrgs[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevOrg {
      if !foundPrevOrgs[i] {
        ctx.org.PrevOrg = append(ctx.org.PrevOrg, w)
      }
    }
  }
  foundScores := make(map[int]bool)
  for i, v := range ctx.org.Score {
    for k, w := range this.Score {
      if w.Same(&v) {
        ctx.org.Score[i] = w
        foundScores[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Score {
      if !foundScores[i] {
        ctx.org.Score = append(ctx.org.Score, w)
      }
    }
  }
  if   !this.DealUid.Same(&ctx.orgDeal.DealUid) {
    ctx.orgDeal.DealUid = this.DealUid
  }
  if   !this.Main.Same(&ctx.orgDeal.Main) {
    ctx.orgDeal.Main = this.Main
  }
  if   !this.Amount.Same(ctx.orgDeal.Amount) {
    ctx.orgDeal.Amount = this.Amount
  }
  if   !this.JointDebtors.Same(ctx.orgDeal.JointDebtors) {
    ctx.orgDeal.JointDebtors = this.JointDebtors
  }
  if   !this.PaymentTerms.Same(ctx.orgDeal.PaymentTerms) {
    ctx.orgDeal.PaymentTerms = this.PaymentTerms
  }
  foundChanges := make(map[int]bool)
  for i, v := range ctx.orgDeal.Change {
    for k, w := range this.Change {
      if w.Same(&v) {
        ctx.orgDeal.Change[i] = w
        foundChanges[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Change {
      if !foundChanges[i] {
        ctx.orgDeal.Change = append(ctx.orgDeal.Change, w)
      }
    }
  }
  if   !this.FundDate.Same(ctx.orgDeal.FundDate) {
    ctx.orgDeal.FundDate = this.FundDate
  }
  foundArrears := make(map[int]bool)
  for i, v := range ctx.orgDeal.Arrear {
    for k, w := range this.Arrear {
      if w.Same(&v) {
        ctx.orgDeal.Arrear[i] = w
        foundArrears[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Arrear {
      if !foundArrears[i] {
        ctx.orgDeal.Arrear = append(ctx.orgDeal.Arrear, w)
      }
    }
  }
  foundDueArrears := make(map[int]bool)
  for i, v := range ctx.orgDeal.DueArrear {
    for k, w := range this.DueArrear {
      if w.Same(&v) {
        ctx.orgDeal.DueArrear[i] = w
        foundDueArrears[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.DueArrear {
      if !foundDueArrears[i] {
        ctx.orgDeal.DueArrear = append(ctx.orgDeal.DueArrear, w)
      }
    }
  }
  foundPastdueArrears := make(map[int]bool)
  for i, v := range ctx.orgDeal.PastdueArrear {
    for k, w := range this.PastdueArrear {
      if w.Same(&v) {
        ctx.orgDeal.PastdueArrear[i] = w
        foundPastdueArrears[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PastdueArrear {
      if !foundPastdueArrears[i] {
        ctx.orgDeal.PastdueArrear = append(ctx.orgDeal.PastdueArrear, w)
      }
    }
  }
  foundPaymentss := make(map[int]bool)
  for i, v := range ctx.orgDeal.Payments {
    for k, w := range this.Payments {
      if w.Same(&v) {
        ctx.orgDeal.Payments[i] = w
        foundPaymentss[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Payments {
      if !foundPaymentss[i] {
        ctx.orgDeal.Payments = append(ctx.orgDeal.Payments, w)
      }
    }
  }
  if   !this.MonthlyPayment.Same(ctx.orgDeal.MonthlyPayment) {
    ctx.orgDeal.MonthlyPayment = this.MonthlyPayment
  }
  if   !this.SourceNonMonetaryObligation.Same(ctx.orgDeal.SourceNonMonetaryObligation) {
    ctx.orgDeal.SourceNonMonetaryObligation = this.SourceNonMonetaryObligation
  }
  if   !this.SubjectNonMonetaryObligation.Same(ctx.orgDeal.SubjectNonMonetaryObligation) {
    ctx.orgDeal.SubjectNonMonetaryObligation = this.SubjectNonMonetaryObligation
  }
  foundCollaterals := make(map[int]bool)
  for i, v := range ctx.orgDeal.Collateral {
    for k, w := range this.Collateral {
      if w.Same(&v) {
        ctx.orgDeal.Collateral[i] = w
        foundCollaterals[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Collateral {
      if !foundCollaterals[i] {
        ctx.orgDeal.Collateral = append(ctx.orgDeal.Collateral, w)
      }
    }
  }
  foundWarrantys := make(map[int]bool)
  for i, v := range ctx.orgDeal.Warranty {
    for k, w := range this.Warranty {
      if w.Same(&v) {
        ctx.orgDeal.Warranty[i] = w
        foundWarrantys[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Warranty {
      if !foundWarrantys[i] {
        ctx.orgDeal.Warranty = append(ctx.orgDeal.Warranty, w)
      }
    }
  }
  foundIndepGuarantees := make(map[int]bool)
  for i, v := range ctx.orgDeal.IndepGuarantee {
    for k, w := range this.IndepGuarantee {
      if w.Same(&v) {
        ctx.orgDeal.IndepGuarantee[i] = w
        foundIndepGuarantees[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.IndepGuarantee {
      if !foundIndepGuarantees[i] {
        ctx.orgDeal.IndepGuarantee = append(ctx.orgDeal.IndepGuarantee, w)
      }
    }
  }
  foundCollateralInsurances := make(map[int]bool)
  for i, v := range ctx.orgDeal.CollateralInsurance {
    for k, w := range this.CollateralInsurance {
      if w.Same(&v) {
        ctx.orgDeal.CollateralInsurance[i] = w
        foundCollateralInsurances[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.CollateralInsurance {
      if !foundCollateralInsurances[i] {
        ctx.orgDeal.CollateralInsurance = append(ctx.orgDeal.CollateralInsurance, w)
      }
    }
  }
  foundSettlements := make(map[int]bool)
  for i, v := range ctx.orgDeal.Settlement {
    for k, w := range this.Settlement {
      if w.Same(&v) {
        ctx.orgDeal.Settlement[i] = w
        foundSettlements[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Settlement {
      if !foundSettlements[i] {
        ctx.orgDeal.Settlement = append(ctx.orgDeal.Settlement, w)
      }
    }
  }
  if   !this.Repayment.Same(ctx.orgDeal.Repayment) {
    ctx.orgDeal.Repayment = this.Repayment
  }
  if   !this.Termination.Same(ctx.orgDeal.Termination) {
    ctx.orgDeal.Termination = this.Termination
  }
  foundLawsuits := make(map[int]bool)
  for i, v := range ctx.orgDeal.Lawsuit {
    for k, w := range this.Lawsuit {
      if w.Same(&v) {
        ctx.orgDeal.Lawsuit[i] = w
        foundLawsuits[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Lawsuit {
      if !foundLawsuits[i] {
        ctx.orgDeal.Lawsuit = append(ctx.orgDeal.Lawsuit, w)
      }
    }
  }
  if   !this.SourceBankruptcy.Same(ctx.orgDeal.SourceBankruptcy) {
    ctx.orgDeal.SourceBankruptcy = this.SourceBankruptcy
  }
  if   !this.SourceLiquidation.Same(ctx.orgDeal.SourceLiquidation) {
    ctx.orgDeal.SourceLiquidation = this.SourceLiquidation
  }
  foundCollection := -1
  for i, v := range ctx.org.Collection {
    if this.Collection.Same(&v) {
      ctx.org.Collection[i] = *this.Collection
      foundCollection = i
    }
  }
  if ctx.op != "D" && foundCollection == -1 {
    ctx.org.Collection = append(ctx.org.Collection, *this.Collection)
  }
  foundOrgCustomers := make(map[int]bool)
  for i, v := range ctx.org.OrgCustomer {
    for k, w := range this.OrgCustomer {
      if w.Same(&v) {
        ctx.org.OrgCustomer[i] = w
        foundOrgCustomers[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.OrgCustomer {
      if !foundOrgCustomers[i] {
        ctx.org.OrgCustomer = append(ctx.org.OrgCustomer, w)
      }
    }
  }
  foundPersonCustomers := make(map[int]bool)
  for i, v := range ctx.org.PersonCustomer {
    for k, w := range this.PersonCustomer {
      if w.Same(&v) {
        ctx.org.PersonCustomer[i] = w
        foundPersonCustomers[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PersonCustomer {
      if !foundPersonCustomers[i] {
        ctx.org.PersonCustomer = append(ctx.org.PersonCustomer, w)
      }
    }
  }
  if   !this.OrgAcquirerInfo.Same(ctx.orgDeal.OrgAcquirerInfo) {
    ctx.orgDeal.OrgAcquirerInfo = this.OrgAcquirerInfo
  }
  if   !this.PersonAcquirerInfo.Same(ctx.orgDeal.PersonAcquirerInfo) {
    ctx.orgDeal.PersonAcquirerInfo = this.PersonAcquirerInfo
  }
  if   !this.ServiceOrg.Same(ctx.orgDeal.ServiceOrg) {
    ctx.orgDeal.ServiceOrg = this.ServiceOrg
  }
  if   !this.Accounting.Same(ctx.orgDeal.Accounting) {
    ctx.orgDeal.Accounting = this.Accounting
  }
  foundApplications := make(map[int]bool)
  for i, v := range ctx.org.Application {
    for k, w := range this.Application {
      if w.Same(&v) {
        ctx.org.Application[i] = w
        foundApplications[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Application {
      if !foundApplications[i] {
        ctx.org.Application = append(ctx.org.Application, w)
      }
    }
  }
  foundParticipations := make(map[int]bool)
  for i, v := range ctx.org.Participation {
    for k, w := range this.Participation {
      if w.Same(&v) {
        ctx.org.Participation[i] = w
        foundParticipations[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Participation {
      if !foundParticipations[i] {
        ctx.org.Participation = append(ctx.org.Participation, w)
      }
    }
  }
  foundAppRejects := make(map[int]bool)
  for i, v := range ctx.org.AppReject {
    for k, w := range this.AppReject {
      if w.Same(&v) {
        ctx.org.AppReject[i] = w
        foundAppRejects[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.AppReject {
      if !foundAppRejects[i] {
        ctx.org.AppReject = append(ctx.org.AppReject, w)
      }
    }
  }
  err := saveSubjectRole(*this.SubjectRole, ctx)
  if nil != err {
    return err
  }
  return nil
}

func (this *XsImportOrgType) Validate(ctx *ParseContext) error {
  if err := XsOrgNameTypeFuncValidate(&this.Name, ctx); err != nil { return err }
  if err := XsOrgAddressTypeFuncValidate(&this.Address, ctx); err != nil { return err }
  if err := XsOrgRegNumTypeFuncValidate(&this.RegNum, ctx); err != nil { return err }
  if this.Tax != nil {
    if err := XsOrgTaxTypeFuncValidate(this.Tax, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgChange {
    if err := XsOrgChangeTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Bankruptcy {
    if err := XsBankruptcyTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Fulfillment {
    if err := XsOrgFulfillmentTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Score {
    if err := XsScoreTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if this.Amount != nil {
    if err := XsDealAmountTypeFuncValidate(this.Amount, ctx); err != nil { return err }
  }
  if this.JointDebtors != nil {
    if err := XsJointDebtorsTypeFuncValidate(this.JointDebtors, ctx); err != nil { return err }
  }
  if this.PaymentTerms != nil {
    if err := XsPaymentTermsTypeFuncValidate(this.PaymentTerms, ctx); err != nil { return err }
  }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncValidate(this.FundDate, ctx); err != nil { return err }
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
  if this.Collection != nil {
    if err := XsCollectionTypeFuncValidate(this.Collection, ctx); err != nil { return err }
  }
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncValidate(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncValidate(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncValidate(this.CommissionerSource, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgCustomer {
    if err := XsOrgCustomerTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PersonCustomer {
    if err := XsPersonCustomerTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncValidate(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncValidate(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  if this.ServiceOrg != nil {
    if err := XsServiceOrgTypeFuncValidate(this.ServiceOrg, ctx); err != nil { return err }
  }
  if this.Accounting != nil {
    if err := XsAccountingTypeFuncValidate(this.Accounting, ctx); err != nil { return err }
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
  return nil
}

// Кредитная история физического лица
type XsPersonCreditHistoryTypeFunc func(this *XsPersonCreditHistoryType, ctx *ParseContext) error

var XsPersonCreditHistoryTypeFuncInit XsPersonCreditHistoryTypeFunc = (*XsPersonCreditHistoryType).Init

var XsPersonCreditHistoryTypeFuncLoad XsPersonCreditHistoryTypeFunc = (*XsPersonCreditHistoryType).Load

var XsPersonCreditHistoryTypeFuncValidate XsPersonCreditHistoryTypeFunc = (*XsPersonCreditHistoryType).Validate

func (this *XsPersonCreditHistoryType) Init(ctx *ParseContext) error {
  if err := XsPersonNameTypeFuncInit(&this.Name, ctx); err != nil { return err }
  for _, v := range (*this).PrevName {
    if err := XsPersonPrevNameTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsBirthInfoTypeFuncInit(&this.BirthInfo, ctx); err != nil { return err }
  for _, v := range (*this).Id {
    if err := XsPersonIdTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevId {
    if err := XsPersonPrevIdTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Tax {
    if err := XsTaxTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.SocialId != nil {
    if err := XsSocialIdTypeFuncInit(this.SocialId, ctx); err != nil { return err }
  }
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
  for _, v := range (*this).Rating {
    if err := XsRatingTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Score {
    if err := XsScoreTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
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
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncInit(this.FundDate, ctx); err != nil { return err }
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
  for _, v := range (*this).AvgPayment {
    if err := XsAvgPaymentTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.SourceBankruptcy != nil {
    if err := XsSourceBankruptcyTypeFuncInit(this.SourceBankruptcy, ctx); err != nil { return err }
  }
  if this.SourceLiquidation != nil {
    if err := XsSourceLiquidationTypeFuncInit(this.SourceLiquidation, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collection {
    if err := XsCollectionTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).CustomerRequest {
    if err := XsCustomerRequestTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).StopInfo {
    if err := XsStopInfoTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncInit(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncInit(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncInit(this.CommissionerSource, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgCustomer {
    if err := XsOrgCustomerTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PersonCustomer {
    if err := XsPersonCustomerTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncInit(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncInit(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  if this.ServiceOrg != nil {
    if err := XsServiceOrgTypeFuncInit(this.ServiceOrg, ctx); err != nil { return err }
  }
  if this.Accounting != nil {
    if err := XsAccountingTypeFuncInit(this.Accounting, ctx); err != nil { return err }
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
  return nil
}

func (this *XsPersonCreditHistoryType) Load(ctx *ParseContext) error {
  if err := XsPersonNameTypeFuncLoad(&this.Name, ctx); err != nil { return err }
  for _, v := range (*this).PrevName {
    if err := XsPersonPrevNameTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if err := XsBirthInfoTypeFuncLoad(&this.BirthInfo, ctx); err != nil { return err }
  for _, v := range (*this).Id {
    if err := XsPersonIdTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevId {
    if err := XsPersonPrevIdTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Tax {
    if err := XsTaxTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if this.SocialId != nil {
    if err := XsSocialIdTypeFuncLoad(this.SocialId, ctx); err != nil { return err }
  }
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
  for _, v := range (*this).Rating {
    if err := XsRatingTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Score {
    if err := XsScoreTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if err := XsDealUidTypeFuncLoad(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncLoad(&this.Main, ctx); err != nil { return err }
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
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncLoad(this.FundDate, ctx); err != nil { return err }
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
  for _, v := range (*this).AvgPayment {
    if err := XsAvgPaymentTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if this.SourceBankruptcy != nil {
    if err := XsSourceBankruptcyTypeFuncLoad(this.SourceBankruptcy, ctx); err != nil { return err }
  }
  if this.SourceLiquidation != nil {
    if err := XsSourceLiquidationTypeFuncLoad(this.SourceLiquidation, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collection {
    if err := XsCollectionTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).CustomerRequest {
    if err := XsCustomerRequestTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).StopInfo {
    if err := XsStopInfoTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncLoad(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncLoad(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncLoad(this.CommissionerSource, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgCustomer {
    if err := XsOrgCustomerTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PersonCustomer {
    if err := XsPersonCustomerTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncLoad(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncLoad(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  if this.ServiceOrg != nil {
    if err := XsServiceOrgTypeFuncLoad(this.ServiceOrg, ctx); err != nil { return err }
  }
  if this.Accounting != nil {
    if err := XsAccountingTypeFuncLoad(this.Accounting, ctx); err != nil { return err }
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
  return nil
}

func (this *XsPersonCreditHistoryType) Validate(ctx *ParseContext) error {
  if err := XsPersonNameTypeFuncValidate(&this.Name, ctx); err != nil { return err }
  for _, v := range (*this).PrevName {
    if err := XsPersonPrevNameTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsBirthInfoTypeFuncValidate(&this.BirthInfo, ctx); err != nil { return err }
  for _, v := range (*this).Id {
    if err := XsPersonIdTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevId {
    if err := XsPersonPrevIdTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Tax {
    if err := XsTaxTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.SocialId != nil {
    if err := XsSocialIdTypeFuncValidate(this.SocialId, ctx); err != nil { return err }
  }
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
  for _, v := range (*this).Rating {
    if err := XsRatingTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Score {
    if err := XsScoreTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsPersonDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
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
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncValidate(this.FundDate, ctx); err != nil { return err }
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
  for _, v := range (*this).AvgPayment {
    if err := XsAvgPaymentTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.SourceBankruptcy != nil {
    if err := XsSourceBankruptcyTypeFuncValidate(this.SourceBankruptcy, ctx); err != nil { return err }
  }
  if this.SourceLiquidation != nil {
    if err := XsSourceLiquidationTypeFuncValidate(this.SourceLiquidation, ctx); err != nil { return err }
  }
  for _, v := range (*this).Collection {
    if err := XsCollectionTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).CustomerRequest {
    if err := XsCustomerRequestTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).StopInfo {
    if err := XsStopInfoTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncValidate(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncValidate(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncValidate(this.CommissionerSource, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgCustomer {
    if err := XsOrgCustomerTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PersonCustomer {
    if err := XsPersonCustomerTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncValidate(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncValidate(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  if this.ServiceOrg != nil {
    if err := XsServiceOrgTypeFuncValidate(this.ServiceOrg, ctx); err != nil { return err }
  }
  if this.Accounting != nil {
    if err := XsAccountingTypeFuncValidate(this.Accounting, ctx); err != nil { return err }
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
  return nil
}

// Блок 1. Имя
type XsPersonNameTypeFunc func(this *XsPersonNameType, ctx *ParseContext) error

var XsPersonNameTypeFuncInit XsPersonNameTypeFunc = (*XsPersonNameType).Init

var XsPersonNameTypeFuncLoad XsPersonNameTypeFunc = (*XsPersonNameType).Load

var XsPersonNameTypeFuncValidate XsPersonNameTypeFunc = (*XsPersonNameType).Validate

func (this *XsPersonNameType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPersonNameType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPersonNameType) Validate(ctx *ParseContext) error {
  return nil
}

type XsPersonNameTypeSameFunc func(this *XsPersonNameType, other *XsPersonNameType) bool

var XsPersonNameTypeFuncSame XsPersonNameTypeSameFunc = (*XsPersonNameType).Same

func (this *XsPersonNameType) Same(other *XsPersonNameType) bool {
  return false
}
// Блок 2. Предыдущее имя
type XsPersonPrevNameTypeFunc func(this *XsPersonPrevNameType, ctx *ParseContext) error

var XsPersonPrevNameTypeFuncInit XsPersonPrevNameTypeFunc = (*XsPersonPrevNameType).Init

var XsPersonPrevNameTypeFuncLoad XsPersonPrevNameTypeFunc = (*XsPersonPrevNameType).Load

var XsPersonPrevNameTypeFuncValidate XsPersonPrevNameTypeFunc = (*XsPersonPrevNameType).Validate

func (this *XsPersonPrevNameType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPersonPrevNameType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPersonPrevNameType) Validate(ctx *ParseContext) error {
  return nil
}

type XsPersonPrevNameTypeSameFunc func(this *XsPersonPrevNameType, other *XsPersonPrevNameType) bool

var XsPersonPrevNameTypeFuncSame XsPersonPrevNameTypeSameFunc = (*XsPersonPrevNameType).Same

func (this *XsPersonPrevNameType) Same(other *XsPersonPrevNameType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// Блок 3. Дата и место рождения
type XsBirthInfoTypeFunc func(this *XsBirthInfoType, ctx *ParseContext) error

var XsBirthInfoTypeFuncInit XsBirthInfoTypeFunc = (*XsBirthInfoType).Init

var XsBirthInfoTypeFuncLoad XsBirthInfoTypeFunc = (*XsBirthInfoType).Load

var XsBirthInfoTypeFuncValidate XsBirthInfoTypeFunc = (*XsBirthInfoType).Validate

func (this *XsBirthInfoType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsBirthInfoType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsBirthInfoType) Validate(ctx *ParseContext) error {
  return nil
}

type XsBirthInfoTypeSameFunc func(this *XsBirthInfoType, other *XsBirthInfoType) bool

var XsBirthInfoTypeFuncSame XsBirthInfoTypeSameFunc = (*XsBirthInfoType).Same

func (this *XsBirthInfoType) Same(other *XsBirthInfoType) bool {
  return false
}
// Блок 4. Документ, удостоверяющий личность; Блок 5. Документ, ранее удостоверявший личность
type XsPersonIdTypeFunc func(this *XsPersonIdType, ctx *ParseContext) error

var XsPersonIdTypeFuncInit XsPersonIdTypeFunc = (*XsPersonIdType).Init

var XsPersonIdTypeFuncLoad XsPersonIdTypeFunc = (*XsPersonIdType).Load

var XsPersonIdTypeFuncValidate XsPersonIdTypeFunc = (*XsPersonIdType).Validate

func (this *XsPersonIdType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPersonIdType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPersonIdType) Validate(ctx *ParseContext) error {
  return nil
}

type XsPersonIdTypeSameFunc func(this *XsPersonIdType, other *XsPersonIdType) bool

var XsPersonIdTypeFuncSame XsPersonIdTypeSameFunc = (*XsPersonIdType).Same

func (this *XsPersonIdType) Same(other *XsPersonIdType) bool {
  if other == nil || this == nil { return false }
  if this.IdCode != other.IdCode { return false }
  return true
}
// Блок 5. Документ, ранее удостоверявший личность
type XsPersonPrevIdTypeFunc func(this *XsPersonPrevIdType, ctx *ParseContext) error

var XsPersonPrevIdTypeFuncInit XsPersonPrevIdTypeFunc = (*XsPersonPrevIdType).Init

var XsPersonPrevIdTypeFuncLoad XsPersonPrevIdTypeFunc = (*XsPersonPrevIdType).Load

var XsPersonPrevIdTypeFuncValidate XsPersonPrevIdTypeFunc = (*XsPersonPrevIdType).Validate

func (this *XsPersonPrevIdType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPersonPrevIdType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPersonPrevIdType) Validate(ctx *ParseContext) error {
  return nil
}

type XsPersonPrevIdTypeSameFunc func(this *XsPersonPrevIdType, other *XsPersonPrevIdType) bool

var XsPersonPrevIdTypeFuncSame XsPersonPrevIdTypeSameFunc = (*XsPersonPrevIdType).Same

func (this *XsPersonPrevIdType) Same(other *XsPersonPrevIdType) bool {
  if other == nil || this == nil { return false }
  if this.IdCode == nil && other.IdCode != nil || this.IdCode != nil && other.IdCode == nil || (this.IdCode != nil && other.IdCode != nil && *this.IdCode != *other.IdCode) { return false }
  if this.IdSeries == nil && other.IdSeries != nil || this.IdSeries != nil && other.IdSeries == nil || (this.IdSeries != nil && other.IdSeries != nil && *this.IdSeries != *other.IdSeries) { return false }
  if this.IdNum == nil && other.IdNum != nil || this.IdNum != nil && other.IdNum == nil || (this.IdNum != nil && other.IdNum != nil && *this.IdNum != *other.IdNum) { return false }
  return true
}
// Блок 6. (юр Блок 6) Номер налогоплательщика и регистрационный номер
type XsTaxTypeFunc func(this *XsTaxType, ctx *ParseContext) error

var XsTaxTypeFuncInit XsTaxTypeFunc = (*XsTaxType).Init

var XsTaxTypeFuncLoad XsTaxTypeFunc = (*XsTaxType).Load

var XsTaxTypeFuncValidate XsTaxTypeFunc = (*XsTaxType).Validate

func (this *XsTaxType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsTaxType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsTaxType) Validate(ctx *ParseContext) error {
  return nil
}

type XsTaxTypeSameFunc func(this *XsTaxType, other *XsTaxType) bool

var XsTaxTypeFuncSame XsTaxTypeSameFunc = (*XsTaxType).Same

func (this *XsTaxType) Same(other *XsTaxType) bool {
  if other == nil || this == nil { return false }
  if this.RegNum == nil && other.RegNum != nil || this.RegNum != nil && other.RegNum == nil || (this.RegNum != nil && other.RegNum != nil && *this.RegNum != *other.RegNum) { return false }
  if this.TaxNum == nil && other.TaxNum != nil || this.TaxNum != nil && other.TaxNum == nil || (this.TaxNum != nil && other.TaxNum != nil && *this.TaxNum != *other.TaxNum) { return false }
  return true
}
// Блок 7. СНИЛС
type XsSocialIdTypeFunc func(this *XsSocialIdType, ctx *ParseContext) error

var XsSocialIdTypeFuncInit XsSocialIdTypeFunc = (*XsSocialIdType).Init

var XsSocialIdTypeFuncLoad XsSocialIdTypeFunc = (*XsSocialIdType).Load

var XsSocialIdTypeFuncValidate XsSocialIdTypeFunc = (*XsSocialIdType).Validate

func (this *XsSocialIdType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsSocialIdType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsSocialIdType) Validate(ctx *ParseContext) error {
  return nil
}

type XsSocialIdTypeSameFunc func(this *XsSocialIdType, other *XsSocialIdType) bool

var XsSocialIdTypeFuncSame XsSocialIdTypeSameFunc = (*XsSocialIdType).Same

func (this *XsSocialIdType) Same(other *XsSocialIdType) bool {
  if other == nil || this == nil { return false }
  if this.Id != other.Id { return false }
  return true
}
// Блок 8. Регистрация физического лица по месту жительства или пребывания
type XsAddressTypeFunc func(this *XsAddressType, ctx *ParseContext) error

var XsAddressTypeFuncInit XsAddressTypeFunc = (*XsAddressType).Init

var XsAddressTypeFuncLoad XsAddressTypeFunc = (*XsAddressType).Load

var XsAddressTypeFuncValidate XsAddressTypeFunc = (*XsAddressType).Validate

func (this *XsAddressType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsAddressType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsAddressType) Validate(ctx *ParseContext) error {
  return nil
}

type XsAddressTypeSameFunc func(this *XsAddressType, other *XsAddressType) bool

var XsAddressTypeFuncSame XsAddressTypeSameFunc = (*XsAddressType).Same

func (this *XsAddressType) Same(other *XsAddressType) bool {
  return false
}
// Блок 9. Фактическое место жительства
type XsFactAddressTypeFunc func(this *XsFactAddressType, ctx *ParseContext) error

var XsFactAddressTypeFuncInit XsFactAddressTypeFunc = (*XsFactAddressType).Init

var XsFactAddressTypeFuncLoad XsFactAddressTypeFunc = (*XsFactAddressType).Load

var XsFactAddressTypeFuncValidate XsFactAddressTypeFunc = (*XsFactAddressType).Validate

func (this *XsFactAddressType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsFactAddressType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsFactAddressType) Validate(ctx *ParseContext) error {
  return nil
}

type XsFactAddressTypeSameFunc func(this *XsFactAddressType, other *XsFactAddressType) bool

var XsFactAddressTypeFuncSame XsFactAddressTypeSameFunc = (*XsFactAddressType).Same

func (this *XsFactAddressType) Same(other *XsFactAddressType) bool {
  return false
}
// Блок 10. Контактные данные
type XsContactTypeFunc func(this *XsContactType, ctx *ParseContext) error

var XsContactTypeFuncInit XsContactTypeFunc = (*XsContactType).Init

var XsContactTypeFuncLoad XsContactTypeFunc = (*XsContactType).Load

var XsContactTypeFuncValidate XsContactTypeFunc = (*XsContactType).Validate

func (this *XsContactType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsContactType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsContactType) Validate(ctx *ParseContext) error {
  return nil
}

type XsContactTypeSameFunc func(this *XsContactType, other *XsContactType) bool

var XsContactTypeFuncSame XsContactTypeSameFunc = (*XsContactType).Same

func (this *XsContactType) Same(other *XsContactType) bool {
  if other == nil || this == nil { return false }
  if this.Phone == nil && other.Phone != nil || this.Phone != nil && other.Phone == nil || (this.Phone != nil && other.Phone != nil && *this.Phone != *other.Phone) { return false }
  if this.Email == nil && other.Email != nil || this.Email != nil && other.Email == nil || (this.Email != nil && other.Email != nil && *this.Email != *other.Email) { return false }
  return true
}
// Блок 11. Государственная регистрация в качестве индивидуального предпринимателя
type XsSoleProprietorTypeFunc func(this *XsSoleProprietorType, ctx *ParseContext) error

var XsSoleProprietorTypeFuncInit XsSoleProprietorTypeFunc = (*XsSoleProprietorType).Init

var XsSoleProprietorTypeFuncLoad XsSoleProprietorTypeFunc = (*XsSoleProprietorType).Load

var XsSoleProprietorTypeFuncValidate XsSoleProprietorTypeFunc = (*XsSoleProprietorType).Validate

func (this *XsSoleProprietorType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsSoleProprietorType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsSoleProprietorType) Validate(ctx *ParseContext) error {
  return nil
}

type XsSoleProprietorTypeSameFunc func(this *XsSoleProprietorType, other *XsSoleProprietorType) bool

var XsSoleProprietorTypeFuncSame XsSoleProprietorTypeSameFunc = (*XsSoleProprietorType).Same

func (this *XsSoleProprietorType) Same(other *XsSoleProprietorType) bool {
  if other == nil || this == nil { return false }
  if this.RegNum == nil && other.RegNum != nil || this.RegNum != nil && other.RegNum == nil || (this.RegNum != nil && other.RegNum != nil && *this.RegNum != *other.RegNum) { return false }
  return true
}
// Блок 12. Сведения о дееспособности
type XsLegalCapacityTypeFunc func(this *XsLegalCapacityType, ctx *ParseContext) error

var XsLegalCapacityTypeFuncInit XsLegalCapacityTypeFunc = (*XsLegalCapacityType).Init

var XsLegalCapacityTypeFuncLoad XsLegalCapacityTypeFunc = (*XsLegalCapacityType).Load

var XsLegalCapacityTypeFuncValidate XsLegalCapacityTypeFunc = (*XsLegalCapacityType).Validate

func (this *XsLegalCapacityType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsLegalCapacityType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsLegalCapacityType) Validate(ctx *ParseContext) error {
  return nil
}

type XsLegalCapacityTypeSameFunc func(this *XsLegalCapacityType, other *XsLegalCapacityType) bool

var XsLegalCapacityTypeFuncSame XsLegalCapacityTypeSameFunc = (*XsLegalCapacityType).Same

func (this *XsLegalCapacityType) Same(other *XsLegalCapacityType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// Блок 13. (юр Блок 6) Сведения по делу о несостоятельности (банкротстве)
type XsBankruptcyTypeFunc func(this *XsBankruptcyType, ctx *ParseContext) error

var XsBankruptcyTypeFuncInit XsBankruptcyTypeFunc = (*XsBankruptcyType).Init

var XsBankruptcyTypeFuncLoad XsBankruptcyTypeFunc = (*XsBankruptcyType).Load

var XsBankruptcyTypeFuncValidate XsBankruptcyTypeFunc = (*XsBankruptcyType).Validate

func (this *XsBankruptcyType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsBankruptcyType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsBankruptcyType) Validate(ctx *ParseContext) error {
  return nil
}

type XsBankruptcyTypeSameFunc func(this *XsBankruptcyType, other *XsBankruptcyType) bool

var XsBankruptcyTypeFuncSame XsBankruptcyTypeSameFunc = (*XsBankruptcyType).Same

func (this *XsBankruptcyType) Same(other *XsBankruptcyType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// Блок 14. Сведения о завершении расчетов с кредиторами и освобождении субъекта от исполнения обязательств в связи с банкротством
type XsFulfillmentTypeFunc func(this *XsFulfillmentType, ctx *ParseContext) error

var XsFulfillmentTypeFuncInit XsFulfillmentTypeFunc = (*XsFulfillmentType).Init

var XsFulfillmentTypeFuncLoad XsFulfillmentTypeFunc = (*XsFulfillmentType).Load

var XsFulfillmentTypeFuncValidate XsFulfillmentTypeFunc = (*XsFulfillmentType).Validate

func (this *XsFulfillmentType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsFulfillmentType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsFulfillmentType) Validate(ctx *ParseContext) error {
  return nil
}

type XsFulfillmentTypeSameFunc func(this *XsFulfillmentType, other *XsFulfillmentType) bool

var XsFulfillmentTypeFuncSame XsFulfillmentTypeSameFunc = (*XsFulfillmentType).Same

func (this *XsFulfillmentType) Same(other *XsFulfillmentType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// Блок 15. Индивидуальный рейтинг субъекта
type XsRatingTypeFunc func(this *XsRatingType, ctx *ParseContext) error

var XsRatingTypeFuncInit XsRatingTypeFunc = (*XsRatingType).Init

var XsRatingTypeFuncLoad XsRatingTypeFunc = (*XsRatingType).Load

var XsRatingTypeFuncValidate XsRatingTypeFunc = (*XsRatingType).Validate

func (this *XsRatingType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsRatingType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsRatingType) Validate(ctx *ParseContext) error {
  return nil
}

type XsRatingTypeSameFunc func(this *XsRatingType, other *XsRatingType) bool

var XsRatingTypeFuncSame XsRatingTypeSameFunc = (*XsRatingType).Same

func (this *XsRatingType) Same(other *XsRatingType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// Блок 16. (юр Блок 9) Кредитная оценка (скоринг)
type XsScoreTypeFunc func(this *XsScoreType, ctx *ParseContext) error

var XsScoreTypeFuncInit XsScoreTypeFunc = (*XsScoreType).Init

var XsScoreTypeFuncLoad XsScoreTypeFunc = (*XsScoreType).Load

var XsScoreTypeFuncValidate XsScoreTypeFunc = (*XsScoreType).Validate

func (this *XsScoreType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsScoreType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsScoreType) Validate(ctx *ParseContext) error {
  return nil
}

type XsScoreTypeSameFunc func(this *XsScoreType, other *XsScoreType) bool

var XsScoreTypeFuncSame XsScoreTypeSameFunc = (*XsScoreType).Same

func (this *XsScoreType) Same(other *XsScoreType) bool {
  if other == nil || this == nil { return false }
  if this.Date != other.Date { return false }
  return true
}
// Блок 17. (юр Блок 10) Уникальный идентификатор договора (сделки)
type XsDealUidTypeFunc func(this *XsDealUidType, ctx *ParseContext) error

var XsDealUidTypeFuncInit XsDealUidTypeFunc = (*XsDealUidType).Init

var XsDealUidTypeFuncLoad XsDealUidTypeFunc = (*XsDealUidType).Load

var XsDealUidTypeFuncValidate XsDealUidTypeFunc = (*XsDealUidType).Validate

func (this *XsDealUidType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsDealUidType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsDealUidType) Validate(ctx *ParseContext) error {
  return nil
}

type XsDealUidTypeSameFunc func(this *XsDealUidType, other *XsDealUidType) bool

var XsDealUidTypeFuncSame XsDealUidTypeSameFunc = (*XsDealUidType).Same

func (this *XsDealUidType) Same(other *XsDealUidType) bool {
  return false
}
// Блок 18. Общие сведения о сделке
type XsPersonDealMainTypeFunc func(this *XsPersonDealMainType, ctx *ParseContext) error

var XsPersonDealMainTypeFuncInit XsPersonDealMainTypeFunc = (*XsPersonDealMainType).Init

var XsPersonDealMainTypeFuncLoad XsPersonDealMainTypeFunc = (*XsPersonDealMainType).Load

var XsPersonDealMainTypeFuncValidate XsPersonDealMainTypeFunc = (*XsPersonDealMainType).Validate

func (this *XsPersonDealMainType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPersonDealMainType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPersonDealMainType) Validate(ctx *ParseContext) error {
  return nil
}

type XsPersonDealMainTypeSameFunc func(this *XsPersonDealMainType, other *XsPersonDealMainType) bool

var XsPersonDealMainTypeFuncSame XsPersonDealMainTypeSameFunc = (*XsPersonDealMainType).Same

func (this *XsPersonDealMainType) Same(other *XsPersonDealMainType) bool {
  return false
}
// Блок 19. (юр Блок 12) Сумма и валюта обязательства
type XsDealAmountTypeFunc func(this *XsDealAmountType, ctx *ParseContext) error

var XsDealAmountTypeFuncInit XsDealAmountTypeFunc = (*XsDealAmountType).Init

var XsDealAmountTypeFuncLoad XsDealAmountTypeFunc = (*XsDealAmountType).Load

var XsDealAmountTypeFuncValidate XsDealAmountTypeFunc = (*XsDealAmountType).Validate

func (this *XsDealAmountType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsDealAmountType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsDealAmountType) Validate(ctx *ParseContext) error {
  return nil
}

type XsDealAmountTypeSameFunc func(this *XsDealAmountType, other *XsDealAmountType) bool

var XsDealAmountTypeFuncSame XsDealAmountTypeSameFunc = (*XsDealAmountType).Same

func (this *XsDealAmountType) Same(other *XsDealAmountType) bool {
  return false
}
// Блок 20. (юр Блок 13) Сведения о солидарных должниках
type XsJointDebtorsTypeFunc func(this *XsJointDebtorsType, ctx *ParseContext) error

var XsJointDebtorsTypeFuncInit XsJointDebtorsTypeFunc = (*XsJointDebtorsType).Init

var XsJointDebtorsTypeFuncLoad XsJointDebtorsTypeFunc = (*XsJointDebtorsType).Load

var XsJointDebtorsTypeFuncValidate XsJointDebtorsTypeFunc = (*XsJointDebtorsType).Validate

func (this *XsJointDebtorsType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsJointDebtorsType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsJointDebtorsType) Validate(ctx *ParseContext) error {
  return nil
}

type XsJointDebtorsTypeSameFunc func(this *XsJointDebtorsType, other *XsJointDebtorsType) bool

var XsJointDebtorsTypeFuncSame XsJointDebtorsTypeSameFunc = (*XsJointDebtorsType).Same

func (this *XsJointDebtorsType) Same(other *XsJointDebtorsType) bool {
  return false
}
// Блок 21. (юр Блок 14) Сведения об условиях платежей
type XsPaymentTermsTypeFunc func(this *XsPaymentTermsType, ctx *ParseContext) error

var XsPaymentTermsTypeFuncInit XsPaymentTermsTypeFunc = (*XsPaymentTermsType).Init

var XsPaymentTermsTypeFuncLoad XsPaymentTermsTypeFunc = (*XsPaymentTermsType).Load

var XsPaymentTermsTypeFuncValidate XsPaymentTermsTypeFunc = (*XsPaymentTermsType).Validate

func (this *XsPaymentTermsType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPaymentTermsType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPaymentTermsType) Validate(ctx *ParseContext) error {
  return nil
}

type XsPaymentTermsTypeSameFunc func(this *XsPaymentTermsType, other *XsPaymentTermsType) bool

var XsPaymentTermsTypeFuncSame XsPaymentTermsTypeSameFunc = (*XsPaymentTermsType).Same

func (this *XsPaymentTermsType) Same(other *XsPaymentTermsType) bool {
  return false
}
// Блок 22. Полная стоимость потребительского кредита (займа)
type XsTotalCostTypeFunc func(this *XsTotalCostType, ctx *ParseContext) error

var XsTotalCostTypeFuncInit XsTotalCostTypeFunc = (*XsTotalCostType).Init

var XsTotalCostTypeFuncLoad XsTotalCostTypeFunc = (*XsTotalCostType).Load

var XsTotalCostTypeFuncValidate XsTotalCostTypeFunc = (*XsTotalCostType).Validate

func (this *XsTotalCostType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsTotalCostType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsTotalCostType) Validate(ctx *ParseContext) error {
  return nil
}

type XsTotalCostTypeSameFunc func(this *XsTotalCostType, other *XsTotalCostType) bool

var XsTotalCostTypeFuncSame XsTotalCostTypeSameFunc = (*XsTotalCostType).Same

func (this *XsTotalCostType) Same(other *XsTotalCostType) bool {
  if other == nil || this == nil { return false }
  if this.Percents != other.Percents { return false }
  if this.Money != other.Money { return false }
  if this.CalcDate != other.CalcDate { return false }
  return true
}
// Блок 23. (юр Блок 15) Сведения об изменении договора
type XsDealChangeTypeFunc func(this *XsDealChangeType, ctx *ParseContext) error

var XsDealChangeTypeFuncInit XsDealChangeTypeFunc = (*XsDealChangeType).Init

var XsDealChangeTypeFuncLoad XsDealChangeTypeFunc = (*XsDealChangeType).Load

var XsDealChangeTypeFuncValidate XsDealChangeTypeFunc = (*XsDealChangeType).Validate

func (this *XsDealChangeType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsDealChangeType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsDealChangeType) Validate(ctx *ParseContext) error {
  return nil
}

type XsDealChangeTypeSameFunc func(this *XsDealChangeType, other *XsDealChangeType) bool

var XsDealChangeTypeFuncSame XsDealChangeTypeSameFunc = (*XsDealChangeType).Same

func (this *XsDealChangeType) Same(other *XsDealChangeType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// Блок 24. (юр Блок 16) Дата передачи финансирования субъекту или возникновения обеспечения исполнения обязательства
type XsFundDateTypeFunc func(this *XsFundDateType, ctx *ParseContext) error

var XsFundDateTypeFuncInit XsFundDateTypeFunc = (*XsFundDateType).Init

var XsFundDateTypeFuncLoad XsFundDateTypeFunc = (*XsFundDateType).Load

var XsFundDateTypeFuncValidate XsFundDateTypeFunc = (*XsFundDateType).Validate

func (this *XsFundDateType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsFundDateType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsFundDateType) Validate(ctx *ParseContext) error {
  return nil
}

type XsFundDateTypeSameFunc func(this *XsFundDateType, other *XsFundDateType) bool

var XsFundDateTypeFuncSame XsFundDateTypeSameFunc = (*XsFundDateType).Same

func (this *XsFundDateType) Same(other *XsFundDateType) bool {
  return false
}
// Блок 25. (юр Блок 17) Сведения о задолженности
type XsArrearTypeFunc func(this *XsArrearType, ctx *ParseContext) error

var XsArrearTypeFuncInit XsArrearTypeFunc = (*XsArrearType).Init

var XsArrearTypeFuncLoad XsArrearTypeFunc = (*XsArrearType).Load

var XsArrearTypeFuncValidate XsArrearTypeFunc = (*XsArrearType).Validate

func (this *XsArrearType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsArrearType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsArrearType) Validate(ctx *ParseContext) error {
  return nil
}

type XsArrearTypeSameFunc func(this *XsArrearType, other *XsArrearType) bool

var XsArrearTypeFuncSame XsArrearTypeSameFunc = (*XsArrearType).Same

func (this *XsArrearType) Same(other *XsArrearType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// Блок 26. (юр Блок 18) Сведения о срочной задолженности
type XsDueArrearTypeFunc func(this *XsDueArrearType, ctx *ParseContext) error

var XsDueArrearTypeFuncInit XsDueArrearTypeFunc = (*XsDueArrearType).Init

var XsDueArrearTypeFuncLoad XsDueArrearTypeFunc = (*XsDueArrearType).Load

var XsDueArrearTypeFuncValidate XsDueArrearTypeFunc = (*XsDueArrearType).Validate

func (this *XsDueArrearType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsDueArrearType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsDueArrearType) Validate(ctx *ParseContext) error {
  return nil
}

type XsDueArrearTypeSameFunc func(this *XsDueArrearType, other *XsDueArrearType) bool

var XsDueArrearTypeFuncSame XsDueArrearTypeSameFunc = (*XsDueArrearType).Same

func (this *XsDueArrearType) Same(other *XsDueArrearType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// Блок 27. (юр Блок 19) Сведения о просроченной задолженности
type XsPastdueArrearTypeFunc func(this *XsPastdueArrearType, ctx *ParseContext) error

var XsPastdueArrearTypeFuncInit XsPastdueArrearTypeFunc = (*XsPastdueArrearType).Init

var XsPastdueArrearTypeFuncLoad XsPastdueArrearTypeFunc = (*XsPastdueArrearType).Load

var XsPastdueArrearTypeFuncValidate XsPastdueArrearTypeFunc = (*XsPastdueArrearType).Validate

func (this *XsPastdueArrearType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPastdueArrearType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPastdueArrearType) Validate(ctx *ParseContext) error {
  return nil
}

type XsPastdueArrearTypeSameFunc func(this *XsPastdueArrearType, other *XsPastdueArrearType) bool

var XsPastdueArrearTypeFuncSame XsPastdueArrearTypeSameFunc = (*XsPastdueArrearType).Same

func (this *XsPastdueArrearType) Same(other *XsPastdueArrearType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// Блок 28. (юр Блок 20) Сведения о внесении платежей
type XsPaymentsTypeFunc func(this *XsPaymentsType, ctx *ParseContext) error

var XsPaymentsTypeFuncInit XsPaymentsTypeFunc = (*XsPaymentsType).Init

var XsPaymentsTypeFuncLoad XsPaymentsTypeFunc = (*XsPaymentsType).Load

var XsPaymentsTypeFuncValidate XsPaymentsTypeFunc = (*XsPaymentsType).Validate

func (this *XsPaymentsType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPaymentsType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPaymentsType) Validate(ctx *ParseContext) error {
  return nil
}

type XsPaymentsTypeSameFunc func(this *XsPaymentsType, other *XsPaymentsType) bool

var XsPaymentsTypeFuncSame XsPaymentsTypeSameFunc = (*XsPaymentsType).Same

func (this *XsPaymentsType) Same(other *XsPaymentsType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// Блок 29. Величина среднемесячного платежа по договору займа (кредита) и дата ее расчета
type XsMonthlyPaymentTypeFunc func(this *XsMonthlyPaymentType, ctx *ParseContext) error

var XsMonthlyPaymentTypeFuncInit XsMonthlyPaymentTypeFunc = (*XsMonthlyPaymentType).Init

var XsMonthlyPaymentTypeFuncLoad XsMonthlyPaymentTypeFunc = (*XsMonthlyPaymentType).Load

var XsMonthlyPaymentTypeFuncValidate XsMonthlyPaymentTypeFunc = (*XsMonthlyPaymentType).Validate

func (this *XsMonthlyPaymentType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsMonthlyPaymentType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsMonthlyPaymentType) Validate(ctx *ParseContext) error {
  return nil
}

type XsMonthlyPaymentTypeSameFunc func(this *XsMonthlyPaymentType, other *XsMonthlyPaymentType) bool

var XsMonthlyPaymentTypeFuncSame XsMonthlyPaymentTypeSameFunc = (*XsMonthlyPaymentType).Same

func (this *XsMonthlyPaymentType) Same(other *XsMonthlyPaymentType) bool {
  if other == nil || this == nil { return false }
  if this.Date != other.Date { return false }
  return true
}
// Блок 30. (юр Блок 21) Сведения о неденежном обязательстве источника
type XsSourceNonMonetaryObligationTypeFunc func(this *XsSourceNonMonetaryObligationType, ctx *ParseContext) error

var XsSourceNonMonetaryObligationTypeFuncInit XsSourceNonMonetaryObligationTypeFunc = (*XsSourceNonMonetaryObligationType).Init

var XsSourceNonMonetaryObligationTypeFuncLoad XsSourceNonMonetaryObligationTypeFunc = (*XsSourceNonMonetaryObligationType).Load

var XsSourceNonMonetaryObligationTypeFuncValidate XsSourceNonMonetaryObligationTypeFunc = (*XsSourceNonMonetaryObligationType).Validate

func (this *XsSourceNonMonetaryObligationType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsSourceNonMonetaryObligationType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsSourceNonMonetaryObligationType) Validate(ctx *ParseContext) error {
  return nil
}

type XsSourceNonMonetaryObligationTypeSameFunc func(this *XsSourceNonMonetaryObligationType, other *XsSourceNonMonetaryObligationType) bool

var XsSourceNonMonetaryObligationTypeFuncSame XsSourceNonMonetaryObligationTypeSameFunc = (*XsSourceNonMonetaryObligationType).Same

func (this *XsSourceNonMonetaryObligationType) Same(other *XsSourceNonMonetaryObligationType) bool {
  return false
}
// Блок 31. (юр Блок 22) Сведения о неденежном обязательстве субъекта
type XsSubjectNonMonetaryObligationTypeFunc func(this *XsSubjectNonMonetaryObligationType, ctx *ParseContext) error

var XsSubjectNonMonetaryObligationTypeFuncInit XsSubjectNonMonetaryObligationTypeFunc = (*XsSubjectNonMonetaryObligationType).Init

var XsSubjectNonMonetaryObligationTypeFuncLoad XsSubjectNonMonetaryObligationTypeFunc = (*XsSubjectNonMonetaryObligationType).Load

var XsSubjectNonMonetaryObligationTypeFuncValidate XsSubjectNonMonetaryObligationTypeFunc = (*XsSubjectNonMonetaryObligationType).Validate

func (this *XsSubjectNonMonetaryObligationType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsSubjectNonMonetaryObligationType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsSubjectNonMonetaryObligationType) Validate(ctx *ParseContext) error {
  return nil
}

type XsSubjectNonMonetaryObligationTypeSameFunc func(this *XsSubjectNonMonetaryObligationType, other *XsSubjectNonMonetaryObligationType) bool

var XsSubjectNonMonetaryObligationTypeFuncSame XsSubjectNonMonetaryObligationTypeSameFunc = (*XsSubjectNonMonetaryObligationType).Same

func (this *XsSubjectNonMonetaryObligationType) Same(other *XsSubjectNonMonetaryObligationType) bool {
  return false
}
// Блок 32. (юр Блок 23) Сведения о залоге
type XsCollateralTypeFunc func(this *XsCollateralType, ctx *ParseContext) error

var XsCollateralTypeFuncInit XsCollateralTypeFunc = (*XsCollateralType).Init

var XsCollateralTypeFuncLoad XsCollateralTypeFunc = (*XsCollateralType).Load

var XsCollateralTypeFuncValidate XsCollateralTypeFunc = (*XsCollateralType).Validate

func (this *XsCollateralType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsCollateralType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsCollateralType) Validate(ctx *ParseContext) error {
  return nil
}

type XsCollateralTypeSameFunc func(this *XsCollateralType, other *XsCollateralType) bool

var XsCollateralTypeFuncSame XsCollateralTypeSameFunc = (*XsCollateralType).Same

func (this *XsCollateralType) Same(other *XsCollateralType) bool {
  if other == nil || this == nil { return false }
  if this.AssetId == nil && other.AssetId != nil || this.AssetId != nil && other.AssetId == nil || (this.AssetId != nil && other.AssetId != nil && *this.AssetId != *other.AssetId) { return false }
  return true
}
// Блок 33. (юр Блок 24) Сведения о поручительстве
type XsWarrantyTypeFunc func(this *XsWarrantyType, ctx *ParseContext) error

var XsWarrantyTypeFuncInit XsWarrantyTypeFunc = (*XsWarrantyType).Init

var XsWarrantyTypeFuncLoad XsWarrantyTypeFunc = (*XsWarrantyType).Load

var XsWarrantyTypeFuncValidate XsWarrantyTypeFunc = (*XsWarrantyType).Validate

func (this *XsWarrantyType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsWarrantyType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsWarrantyType) Validate(ctx *ParseContext) error {
  return nil
}

type XsWarrantyTypeSameFunc func(this *XsWarrantyType, other *XsWarrantyType) bool

var XsWarrantyTypeFuncSame XsWarrantyTypeSameFunc = (*XsWarrantyType).Same

func (this *XsWarrantyType) Same(other *XsWarrantyType) bool {
  if other == nil || this == nil { return false }
  if this.ContractUid == nil && other.ContractUid != nil || this.ContractUid != nil && other.ContractUid == nil || (this.ContractUid != nil && other.ContractUid != nil && *this.ContractUid != *other.ContractUid) { return false }
  return true
}
// Блок 34. (юр Блок 25) Сведения о независимой гарантии
type XsIndepGuaranteeTypeFunc func(this *XsIndepGuaranteeType, ctx *ParseContext) error

var XsIndepGuaranteeTypeFuncInit XsIndepGuaranteeTypeFunc = (*XsIndepGuaranteeType).Init

var XsIndepGuaranteeTypeFuncLoad XsIndepGuaranteeTypeFunc = (*XsIndepGuaranteeType).Load

var XsIndepGuaranteeTypeFuncValidate XsIndepGuaranteeTypeFunc = (*XsIndepGuaranteeType).Validate

func (this *XsIndepGuaranteeType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsIndepGuaranteeType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsIndepGuaranteeType) Validate(ctx *ParseContext) error {
  return nil
}

type XsIndepGuaranteeTypeSameFunc func(this *XsIndepGuaranteeType, other *XsIndepGuaranteeType) bool

var XsIndepGuaranteeTypeFuncSame XsIndepGuaranteeTypeSameFunc = (*XsIndepGuaranteeType).Same

func (this *XsIndepGuaranteeType) Same(other *XsIndepGuaranteeType) bool {
  if other == nil || this == nil { return false }
  if this.ContractUid == nil && other.ContractUid != nil || this.ContractUid != nil && other.ContractUid == nil || (this.ContractUid != nil && other.ContractUid != nil && *this.ContractUid != *other.ContractUid) { return false }
  return true
}
// Блок 35. (юр Блок 26) Сведения о страховании предмета залога
type XsCollateralInsuranceTypeFunc func(this *XsCollateralInsuranceType, ctx *ParseContext) error

var XsCollateralInsuranceTypeFuncInit XsCollateralInsuranceTypeFunc = (*XsCollateralInsuranceType).Init

var XsCollateralInsuranceTypeFuncLoad XsCollateralInsuranceTypeFunc = (*XsCollateralInsuranceType).Load

var XsCollateralInsuranceTypeFuncValidate XsCollateralInsuranceTypeFunc = (*XsCollateralInsuranceType).Validate

func (this *XsCollateralInsuranceType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsCollateralInsuranceType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsCollateralInsuranceType) Validate(ctx *ParseContext) error {
  return nil
}

type XsCollateralInsuranceTypeSameFunc func(this *XsCollateralInsuranceType, other *XsCollateralInsuranceType) bool

var XsCollateralInsuranceTypeFuncSame XsCollateralInsuranceTypeSameFunc = (*XsCollateralInsuranceType).Same

func (this *XsCollateralInsuranceType) Same(other *XsCollateralInsuranceType) bool {
  if other == nil || this == nil { return false }
  if this.AssetId == nil && other.AssetId != nil || this.AssetId != nil && other.AssetId == nil || (this.AssetId != nil && other.AssetId != nil && *this.AssetId != *other.AssetId) { return false }
  return true
}
// Блок 36. (юр Блок 27) Сведения о погашении требований кредитора по обязательству за счет обеспечения
type XsSettlementTypeFunc func(this *XsSettlementType, ctx *ParseContext) error

var XsSettlementTypeFuncInit XsSettlementTypeFunc = (*XsSettlementType).Init

var XsSettlementTypeFuncLoad XsSettlementTypeFunc = (*XsSettlementType).Load

var XsSettlementTypeFuncValidate XsSettlementTypeFunc = (*XsSettlementType).Validate

func (this *XsSettlementType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsSettlementType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsSettlementType) Validate(ctx *ParseContext) error {
  return nil
}

type XsSettlementTypeSameFunc func(this *XsSettlementType, other *XsSettlementType) bool

var XsSettlementTypeFuncSame XsSettlementTypeSameFunc = (*XsSettlementType).Same

func (this *XsSettlementType) Same(other *XsSettlementType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// Блок 37. (юр Блок 28) Сведения о возмещении принципалом гаранту выплаченной суммы
type XsRepaymentTypeFunc func(this *XsRepaymentType, ctx *ParseContext) error

var XsRepaymentTypeFuncInit XsRepaymentTypeFunc = (*XsRepaymentType).Init

var XsRepaymentTypeFuncLoad XsRepaymentTypeFunc = (*XsRepaymentType).Load

var XsRepaymentTypeFuncValidate XsRepaymentTypeFunc = (*XsRepaymentType).Validate

func (this *XsRepaymentType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsRepaymentType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsRepaymentType) Validate(ctx *ParseContext) error {
  return nil
}

type XsRepaymentTypeSameFunc func(this *XsRepaymentType, other *XsRepaymentType) bool

var XsRepaymentTypeFuncSame XsRepaymentTypeSameFunc = (*XsRepaymentType).Same

func (this *XsRepaymentType) Same(other *XsRepaymentType) bool {
  return false
}
// Блок 38. (юр Блок 29) Сведения о прекращении обязательства
type XsTerminationTypeFunc func(this *XsTerminationType, ctx *ParseContext) error

var XsTerminationTypeFuncInit XsTerminationTypeFunc = (*XsTerminationType).Init

var XsTerminationTypeFuncLoad XsTerminationTypeFunc = (*XsTerminationType).Load

var XsTerminationTypeFuncValidate XsTerminationTypeFunc = (*XsTerminationType).Validate

func (this *XsTerminationType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsTerminationType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsTerminationType) Validate(ctx *ParseContext) error {
  return nil
}

type XsTerminationTypeSameFunc func(this *XsTerminationType, other *XsTerminationType) bool

var XsTerminationTypeFuncSame XsTerminationTypeSameFunc = (*XsTerminationType).Same

func (this *XsTerminationType) Same(other *XsTerminationType) bool {
  return false
}
// Блок 39. (юр Блок 30) Сведения о судебном споре или требовании по обязательству
type XsLawsuitTypeFunc func(this *XsLawsuitType, ctx *ParseContext) error

var XsLawsuitTypeFuncInit XsLawsuitTypeFunc = (*XsLawsuitType).Init

var XsLawsuitTypeFuncLoad XsLawsuitTypeFunc = (*XsLawsuitType).Load

var XsLawsuitTypeFuncValidate XsLawsuitTypeFunc = (*XsLawsuitType).Validate

func (this *XsLawsuitType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsLawsuitType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsLawsuitType) Validate(ctx *ParseContext) error {
  return nil
}

type XsLawsuitTypeSameFunc func(this *XsLawsuitType, other *XsLawsuitType) bool

var XsLawsuitTypeFuncSame XsLawsuitTypeSameFunc = (*XsLawsuitType).Same

func (this *XsLawsuitType) Same(other *XsLawsuitType) bool {
  if other == nil || this == nil { return false }
  if this.ActNum == nil && other.ActNum != nil || this.ActNum != nil && other.ActNum == nil || (this.ActNum != nil && other.ActNum != nil && *this.ActNum != *other.ActNum) { return false }
  return true
}
// Блок 40. Сведения квалифицированного бюро о среднемесячных платежах по договору займа (кредита)
type XsAvgPaymentTypeFunc func(this *XsAvgPaymentType, ctx *ParseContext) error

var XsAvgPaymentTypeFuncInit XsAvgPaymentTypeFunc = (*XsAvgPaymentType).Init

var XsAvgPaymentTypeFuncLoad XsAvgPaymentTypeFunc = (*XsAvgPaymentType).Load

var XsAvgPaymentTypeFuncValidate XsAvgPaymentTypeFunc = (*XsAvgPaymentType).Validate

func (this *XsAvgPaymentType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsAvgPaymentType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsAvgPaymentType) Validate(ctx *ParseContext) error {
  return nil
}

type XsAvgPaymentTypeSameFunc func(this *XsAvgPaymentType, other *XsAvgPaymentType) bool

var XsAvgPaymentTypeFuncSame XsAvgPaymentTypeSameFunc = (*XsAvgPaymentType).Same

func (this *XsAvgPaymentType) Same(other *XsAvgPaymentType) bool {
  if other == nil || this == nil { return false }
  if this.Date != other.Date { return false }
  return true
}
// Блок 41. (юр Блок 31) Сведения об обязательстве, если в отношении источника открыто конкурсное производство
type XsSourceBankruptcyTypeFunc func(this *XsSourceBankruptcyType, ctx *ParseContext) error

var XsSourceBankruptcyTypeFuncInit XsSourceBankruptcyTypeFunc = (*XsSourceBankruptcyType).Init

var XsSourceBankruptcyTypeFuncLoad XsSourceBankruptcyTypeFunc = (*XsSourceBankruptcyType).Load

var XsSourceBankruptcyTypeFuncValidate XsSourceBankruptcyTypeFunc = (*XsSourceBankruptcyType).Validate

func (this *XsSourceBankruptcyType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsSourceBankruptcyType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsSourceBankruptcyType) Validate(ctx *ParseContext) error {
  return nil
}

type XsSourceBankruptcyTypeSameFunc func(this *XsSourceBankruptcyType, other *XsSourceBankruptcyType) bool

var XsSourceBankruptcyTypeFuncSame XsSourceBankruptcyTypeSameFunc = (*XsSourceBankruptcyType).Same

func (this *XsSourceBankruptcyType) Same(other *XsSourceBankruptcyType) bool {
  return false
}
// Блок 42. (юр Блок 32) Сведения об обязательстве, если источник находится в процессе ликвидации
type XsSourceLiquidationTypeFunc func(this *XsSourceLiquidationType, ctx *ParseContext) error

var XsSourceLiquidationTypeFuncInit XsSourceLiquidationTypeFunc = (*XsSourceLiquidationType).Init

var XsSourceLiquidationTypeFuncLoad XsSourceLiquidationTypeFunc = (*XsSourceLiquidationType).Load

var XsSourceLiquidationTypeFuncValidate XsSourceLiquidationTypeFunc = (*XsSourceLiquidationType).Validate

func (this *XsSourceLiquidationType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsSourceLiquidationType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsSourceLiquidationType) Validate(ctx *ParseContext) error {
  return nil
}

type XsSourceLiquidationTypeSameFunc func(this *XsSourceLiquidationType, other *XsSourceLiquidationType) bool

var XsSourceLiquidationTypeFuncSame XsSourceLiquidationTypeSameFunc = (*XsSourceLiquidationType).Same

func (this *XsSourceLiquidationType) Same(other *XsSourceLiquidationType) bool {
  return false
}
// Блок 43. (юр Блок 33) Сведения о взыскании долга по алиментам, платы за жилое помещение, коммунальные услуги или услуги связи
type XsCollectionTypeFunc func(this *XsCollectionType, ctx *ParseContext) error

var XsCollectionTypeFuncInit XsCollectionTypeFunc = (*XsCollectionType).Init

var XsCollectionTypeFuncLoad XsCollectionTypeFunc = (*XsCollectionType).Load

var XsCollectionTypeFuncValidate XsCollectionTypeFunc = (*XsCollectionType).Validate

func (this *XsCollectionType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsCollectionType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsCollectionType) Validate(ctx *ParseContext) error {
  return nil
}

type XsCollectionTypeSameFunc func(this *XsCollectionType, other *XsCollectionType) bool

var XsCollectionTypeFuncSame XsCollectionTypeSameFunc = (*XsCollectionType).Same

func (this *XsCollectionType) Same(other *XsCollectionType) bool {
  if other == nil || this == nil { return false }
  if this.ActNum != other.ActNum { return false }
  return true
}
// Блок 44. (юр Блок 34) Сведения о запросе информации пользователем
type XsCustomerRequestTypeFunc func(this *XsCustomerRequestType, ctx *ParseContext) error

var XsCustomerRequestTypeFuncInit XsCustomerRequestTypeFunc = (*XsCustomerRequestType).Init

var XsCustomerRequestTypeFuncLoad XsCustomerRequestTypeFunc = (*XsCustomerRequestType).Load

var XsCustomerRequestTypeFuncValidate XsCustomerRequestTypeFunc = (*XsCustomerRequestType).Validate

func (this *XsCustomerRequestType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsCustomerRequestType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsCustomerRequestType) Validate(ctx *ParseContext) error {
  return nil
}

// Блок 45. (юр Блок 35) Сведения о прекращении передачи информации по обязательству
type XsStopInfoTypeFunc func(this *XsStopInfoType, ctx *ParseContext) error

var XsStopInfoTypeFuncInit XsStopInfoTypeFunc = (*XsStopInfoType).Init

var XsStopInfoTypeFuncLoad XsStopInfoTypeFunc = (*XsStopInfoType).Load

var XsStopInfoTypeFuncValidate XsStopInfoTypeFunc = (*XsStopInfoType).Validate

func (this *XsStopInfoType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsStopInfoType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsStopInfoType) Validate(ctx *ParseContext) error {
  return nil
}

type XsStopInfoTypeSameFunc func(this *XsStopInfoType, other *XsStopInfoType) bool

var XsStopInfoTypeFuncSame XsStopInfoTypeSameFunc = (*XsStopInfoType).Same

func (this *XsStopInfoType) Same(other *XsStopInfoType) bool {
  if other == nil || this == nil { return false }
  if this.Date != other.Date { return false }
  return true
}
// Блок 46. (юр Блок 36) Сведения об источнике – юридическом лице
type XsOrgSourceTypeFunc func(this *XsOrgSourceType, ctx *ParseContext) error

var XsOrgSourceTypeFuncInit XsOrgSourceTypeFunc = (*XsOrgSourceType).Init

var XsOrgSourceTypeFuncLoad XsOrgSourceTypeFunc = (*XsOrgSourceType).Load

var XsOrgSourceTypeFuncValidate XsOrgSourceTypeFunc = (*XsOrgSourceType).Validate

func (this *XsOrgSourceType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsOrgSourceType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsOrgSourceType) Validate(ctx *ParseContext) error {
  return nil
}

// Блок 47. (юр Блок 37) Сведения об источнике – физическом лице
type XsPersonSourceTypeFunc func(this *XsPersonSourceType, ctx *ParseContext) error

var XsPersonSourceTypeFuncInit XsPersonSourceTypeFunc = (*XsPersonSourceType).Init

var XsPersonSourceTypeFuncLoad XsPersonSourceTypeFunc = (*XsPersonSourceType).Load

var XsPersonSourceTypeFuncValidate XsPersonSourceTypeFunc = (*XsPersonSourceType).Validate

func (this *XsPersonSourceType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPersonSourceType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPersonSourceType) Validate(ctx *ParseContext) error {
  return nil
}

// Блок 48. (юр Блок 38) Сведения об источнике – арбитражном управляющем
type XsCommissionerSourceTypeFunc func(this *XsCommissionerSourceType, ctx *ParseContext) error

var XsCommissionerSourceTypeFuncInit XsCommissionerSourceTypeFunc = (*XsCommissionerSourceType).Init

var XsCommissionerSourceTypeFuncLoad XsCommissionerSourceTypeFunc = (*XsCommissionerSourceType).Load

var XsCommissionerSourceTypeFuncValidate XsCommissionerSourceTypeFunc = (*XsCommissionerSourceType).Validate

func (this *XsCommissionerSourceType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsCommissionerSourceType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsCommissionerSourceType) Validate(ctx *ParseContext) error {
  return nil
}

// Блок 49. (юр Блок 39) Сведения о пользователе – юридическом лице
type XsOrgCustomerTypeFunc func(this *XsOrgCustomerType, ctx *ParseContext) error

var XsOrgCustomerTypeFuncInit XsOrgCustomerTypeFunc = (*XsOrgCustomerType).Init

var XsOrgCustomerTypeFuncLoad XsOrgCustomerTypeFunc = (*XsOrgCustomerType).Load

var XsOrgCustomerTypeFuncValidate XsOrgCustomerTypeFunc = (*XsOrgCustomerType).Validate

func (this *XsOrgCustomerType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsOrgCustomerType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsOrgCustomerType) Validate(ctx *ParseContext) error {
  return nil
}

type XsOrgCustomerTypeSameFunc func(this *XsOrgCustomerType, other *XsOrgCustomerType) bool

var XsOrgCustomerTypeFuncSame XsOrgCustomerTypeSameFunc = (*XsOrgCustomerType).Same

func (this *XsOrgCustomerType) Same(other *XsOrgCustomerType) bool {
  if other == nil || this == nil { return false }
  if this.RegNum == nil && other.RegNum != nil || this.RegNum != nil && other.RegNum == nil || (this.RegNum != nil && other.RegNum != nil && *this.RegNum != *other.RegNum) { return false }
  return true
}
// Блок 50. (юр Блок 40) Сведения о пользователе – индивидуальном предпринимателе
type XsPersonCustomerTypeFunc func(this *XsPersonCustomerType, ctx *ParseContext) error

var XsPersonCustomerTypeFuncInit XsPersonCustomerTypeFunc = (*XsPersonCustomerType).Init

var XsPersonCustomerTypeFuncLoad XsPersonCustomerTypeFunc = (*XsPersonCustomerType).Load

var XsPersonCustomerTypeFuncValidate XsPersonCustomerTypeFunc = (*XsPersonCustomerType).Validate

func (this *XsPersonCustomerType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPersonCustomerType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPersonCustomerType) Validate(ctx *ParseContext) error {
  return nil
}

type XsPersonCustomerTypeSameFunc func(this *XsPersonCustomerType, other *XsPersonCustomerType) bool

var XsPersonCustomerTypeFuncSame XsPersonCustomerTypeSameFunc = (*XsPersonCustomerType).Same

func (this *XsPersonCustomerType) Same(other *XsPersonCustomerType) bool {
  if other == nil || this == nil { return false }
  if this.IdCode != other.IdCode { return false }
  if this.IdSeries == nil && other.IdSeries != nil || this.IdSeries != nil && other.IdSeries == nil || (this.IdSeries != nil && other.IdSeries != nil && *this.IdSeries != *other.IdSeries) { return false }
  if this.IdNum != other.IdNum { return false }
  return true
}
// Блок 51. (юр Блок 41) Сведения о приобретателе прав – юридическом лице
type XsOrgAcquirerInfoTypeFunc func(this *XsOrgAcquirerInfoType, ctx *ParseContext) error

var XsOrgAcquirerInfoTypeFuncInit XsOrgAcquirerInfoTypeFunc = (*XsOrgAcquirerInfoType).Init

var XsOrgAcquirerInfoTypeFuncLoad XsOrgAcquirerInfoTypeFunc = (*XsOrgAcquirerInfoType).Load

var XsOrgAcquirerInfoTypeFuncValidate XsOrgAcquirerInfoTypeFunc = (*XsOrgAcquirerInfoType).Validate

func (this *XsOrgAcquirerInfoType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsOrgAcquirerInfoType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsOrgAcquirerInfoType) Validate(ctx *ParseContext) error {
  return nil
}

type XsOrgAcquirerInfoTypeSameFunc func(this *XsOrgAcquirerInfoType, other *XsOrgAcquirerInfoType) bool

var XsOrgAcquirerInfoTypeFuncSame XsOrgAcquirerInfoTypeSameFunc = (*XsOrgAcquirerInfoType).Same

func (this *XsOrgAcquirerInfoType) Same(other *XsOrgAcquirerInfoType) bool {
  return false
}
// Блок 52. (юр Блок 42) Сведения о приобретателе прав – физическом лице
type XsPersonAcquirerInfoTypeFunc func(this *XsPersonAcquirerInfoType, ctx *ParseContext) error

var XsPersonAcquirerInfoTypeFuncInit XsPersonAcquirerInfoTypeFunc = (*XsPersonAcquirerInfoType).Init

var XsPersonAcquirerInfoTypeFuncLoad XsPersonAcquirerInfoTypeFunc = (*XsPersonAcquirerInfoType).Load

var XsPersonAcquirerInfoTypeFuncValidate XsPersonAcquirerInfoTypeFunc = (*XsPersonAcquirerInfoType).Validate

func (this *XsPersonAcquirerInfoType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPersonAcquirerInfoType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPersonAcquirerInfoType) Validate(ctx *ParseContext) error {
  return nil
}

type XsPersonAcquirerInfoTypeSameFunc func(this *XsPersonAcquirerInfoType, other *XsPersonAcquirerInfoType) bool

var XsPersonAcquirerInfoTypeFuncSame XsPersonAcquirerInfoTypeSameFunc = (*XsPersonAcquirerInfoType).Same

func (this *XsPersonAcquirerInfoType) Same(other *XsPersonAcquirerInfoType) bool {
  return false
}
// Блок 53. (юр Блок 43) Сведения об обслуживающей организации
type XsServiceOrgTypeFunc func(this *XsServiceOrgType, ctx *ParseContext) error

var XsServiceOrgTypeFuncInit XsServiceOrgTypeFunc = (*XsServiceOrgType).Init

var XsServiceOrgTypeFuncLoad XsServiceOrgTypeFunc = (*XsServiceOrgType).Load

var XsServiceOrgTypeFuncValidate XsServiceOrgTypeFunc = (*XsServiceOrgType).Validate

func (this *XsServiceOrgType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsServiceOrgType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsServiceOrgType) Validate(ctx *ParseContext) error {
  return nil
}

type XsServiceOrgTypeSameFunc func(this *XsServiceOrgType, other *XsServiceOrgType) bool

var XsServiceOrgTypeFuncSame XsServiceOrgTypeSameFunc = (*XsServiceOrgType).Same

func (this *XsServiceOrgType) Same(other *XsServiceOrgType) bool {
  return false
}
// Блок 54. (юр Блок 44) Сведения об учете обязательства
type XsAccountingTypeFunc func(this *XsAccountingType, ctx *ParseContext) error

var XsAccountingTypeFuncInit XsAccountingTypeFunc = (*XsAccountingType).Init

var XsAccountingTypeFuncLoad XsAccountingTypeFunc = (*XsAccountingType).Load

var XsAccountingTypeFuncValidate XsAccountingTypeFunc = (*XsAccountingType).Validate

func (this *XsAccountingType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsAccountingType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsAccountingType) Validate(ctx *ParseContext) error {
  return nil
}

type XsAccountingTypeSameFunc func(this *XsAccountingType, other *XsAccountingType) bool

var XsAccountingTypeFuncSame XsAccountingTypeSameFunc = (*XsAccountingType).Same

func (this *XsAccountingType) Same(other *XsAccountingType) bool {
  return false
}
// Блок 55. (юр Блок 45) Сведения об обращении субъекта к источнику с предложением совершить сделку
type XsApplicationTypeFunc func(this *XsApplicationType, ctx *ParseContext) error

var XsApplicationTypeFuncInit XsApplicationTypeFunc = (*XsApplicationType).Init

var XsApplicationTypeFuncLoad XsApplicationTypeFunc = (*XsApplicationType).Load

var XsApplicationTypeFuncValidate XsApplicationTypeFunc = (*XsApplicationType).Validate

func (this *XsApplicationType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsApplicationType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsApplicationType) Validate(ctx *ParseContext) error {
  return nil
}

type XsApplicationTypeSameFunc func(this *XsApplicationType, other *XsApplicationType) bool

var XsApplicationTypeFuncSame XsApplicationTypeSameFunc = (*XsApplicationType).Same

func (this *XsApplicationType) Same(other *XsApplicationType) bool {
  if other == nil || this == nil { return false }
  if this.Uid != other.Uid { return false }
  return true
}
// Блок 56. (юр Блок 46) Сведения об участии в обязательстве, по которому формируется кредитная история
type XsParticipationTypeFunc func(this *XsParticipationType, ctx *ParseContext) error

var XsParticipationTypeFuncInit XsParticipationTypeFunc = (*XsParticipationType).Init

var XsParticipationTypeFuncLoad XsParticipationTypeFunc = (*XsParticipationType).Load

var XsParticipationTypeFuncValidate XsParticipationTypeFunc = (*XsParticipationType).Validate

func (this *XsParticipationType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsParticipationType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsParticipationType) Validate(ctx *ParseContext) error {
  return nil
}

type XsParticipationTypeSameFunc func(this *XsParticipationType, other *XsParticipationType) bool

var XsParticipationTypeFuncSame XsParticipationTypeSameFunc = (*XsParticipationType).Same

func (this *XsParticipationType) Same(other *XsParticipationType) bool {
  if other == nil || this == nil { return false }
  if this.DealUid != other.DealUid { return false }
  return true
}
// Блок 57. (юр Блок 47) Сведения об отказе источника от предложения совершить сделку
type XsAppRejectTypeFunc func(this *XsAppRejectType, ctx *ParseContext) error

var XsAppRejectTypeFuncInit XsAppRejectTypeFunc = (*XsAppRejectType).Init

var XsAppRejectTypeFuncLoad XsAppRejectTypeFunc = (*XsAppRejectType).Load

var XsAppRejectTypeFuncValidate XsAppRejectTypeFunc = (*XsAppRejectType).Validate

func (this *XsAppRejectType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsAppRejectType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsAppRejectType) Validate(ctx *ParseContext) error {
  return nil
}

type XsAppRejectTypeSameFunc func(this *XsAppRejectType, other *XsAppRejectType) bool

var XsAppRejectTypeFuncSame XsAppRejectTypeSameFunc = (*XsAppRejectType).Same

func (this *XsAppRejectType) Same(other *XsAppRejectType) bool {
  return false
}
// Титульная часть КИ Субъекта, используется только для поиска Субъекта
type XsPersonSubjectTypeFunc func(this *XsPersonSubjectType, ctx *ParseContext) error

var XsPersonSubjectTypeFuncInit XsPersonSubjectTypeFunc = (*XsPersonSubjectType).Init

var XsPersonSubjectTypeFuncLoad XsPersonSubjectTypeFunc = (*XsPersonSubjectType).Load

var XsPersonSubjectTypeFuncValidate XsPersonSubjectTypeFunc = (*XsPersonSubjectType).Validate

func (this *XsPersonSubjectType) Init(ctx *ParseContext) error {
  if err := XsPersonNameTypeFuncInit(&this.Name, ctx); err != nil { return err }
  for _, v := range (*this).PrevName {
    if err := XsPersonPrevNameTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsBirthInfoTypeFuncInit(&this.BirthInfo, ctx); err != nil { return err }
  for _, v := range (*this).Id {
    if err := XsPersonIdTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevId {
    if err := XsPersonPrevIdTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Tax {
    if err := XsTaxTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.SocialId != nil {
    if err := XsSocialIdTypeFuncInit(this.SocialId, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsPersonSubjectType) Load(ctx *ParseContext) error {
  if   !this.Name.Same(&ctx.person.Name) {
    ctx.person.Name = this.Name
  }
  foundPrevNames := make(map[int]bool)
  for i, v := range ctx.person.PrevName {
    for k, w := range this.PrevName {
      if w.Same(&v) {
        ctx.person.PrevName[i] = w
        foundPrevNames[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevName {
      if !foundPrevNames[i] {
        ctx.person.PrevName = append(ctx.person.PrevName, w)
      }
    }
  }
  if   !this.BirthInfo.Same(&ctx.person.BirthInfo) {
    ctx.person.BirthInfo = this.BirthInfo
  }
  foundIds := make(map[int]bool)
  for i, v := range ctx.person.Id {
    for k, w := range this.Id {
      if w.Same(&v) {
        ctx.person.Id[i] = w
        foundIds[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Id {
      if !foundIds[i] {
        ctx.person.Id = append(ctx.person.Id, w)
      }
    }
  }
  foundPrevIds := make(map[int]bool)
  for i, v := range ctx.person.PrevId {
    for k, w := range this.PrevId {
      if w.Same(&v) {
        ctx.person.PrevId[i] = w
        foundPrevIds[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.PrevId {
      if !foundPrevIds[i] {
        ctx.person.PrevId = append(ctx.person.PrevId, w)
      }
    }
  }
  foundTaxs := make(map[int]bool)
  for i, v := range ctx.person.Tax {
    for k, w := range this.Tax {
      if w.Same(&v) {
        ctx.person.Tax[i] = w
        foundTaxs[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.Tax {
      if !foundTaxs[i] {
        ctx.person.Tax = append(ctx.person.Tax, w)
      }
    }
  }
  if   !this.SocialId.Same(ctx.person.SocialId) {
    ctx.person.SocialId = this.SocialId
  }
  return nil
}

func (this *XsPersonSubjectType) Validate(ctx *ParseContext) error {
  if err := XsPersonNameTypeFuncValidate(&this.Name, ctx); err != nil { return err }
  for _, v := range (*this).PrevName {
    if err := XsPersonPrevNameTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsBirthInfoTypeFuncValidate(&this.BirthInfo, ctx); err != nil { return err }
  for _, v := range (*this).Id {
    if err := XsPersonIdTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevId {
    if err := XsPersonPrevIdTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Tax {
    if err := XsTaxTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.SocialId != nil {
    if err := XsSocialIdTypeFuncValidate(this.SocialId, ctx); err != nil { return err }
  }
  return nil
}

// Базовый общий блок для источника - физического лица
type XsPersonSourceCommonTypeFunc func(this *XsPersonSourceCommonType, ctx *ParseContext) error

var XsPersonSourceCommonTypeFuncInit XsPersonSourceCommonTypeFunc = (*XsPersonSourceCommonType).Init

var XsPersonSourceCommonTypeFuncLoad XsPersonSourceCommonTypeFunc = (*XsPersonSourceCommonType).Load

var XsPersonSourceCommonTypeFuncValidate XsPersonSourceCommonTypeFunc = (*XsPersonSourceCommonType).Validate

func (this *XsPersonSourceCommonType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPersonSourceCommonType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPersonSourceCommonType) Validate(ctx *ParseContext) error {
  return nil
}

// Базовый общий блок для источника - юридического лица
type XsOrgSourceCommonTypeFunc func(this *XsOrgSourceCommonType, ctx *ParseContext) error

var XsOrgSourceCommonTypeFuncInit XsOrgSourceCommonTypeFunc = (*XsOrgSourceCommonType).Init

var XsOrgSourceCommonTypeFuncLoad XsOrgSourceCommonTypeFunc = (*XsOrgSourceCommonType).Load

var XsOrgSourceCommonTypeFuncValidate XsOrgSourceCommonTypeFunc = (*XsOrgSourceCommonType).Validate

func (this *XsOrgSourceCommonType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsOrgSourceCommonType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsOrgSourceCommonType) Validate(ctx *ParseContext) error {
  return nil
}

// Кредитная история юридического лица
type XsOrgCreditHistoryTypeFunc func(this *XsOrgCreditHistoryType, ctx *ParseContext) error

var XsOrgCreditHistoryTypeFuncInit XsOrgCreditHistoryTypeFunc = (*XsOrgCreditHistoryType).Init

var XsOrgCreditHistoryTypeFuncLoad XsOrgCreditHistoryTypeFunc = (*XsOrgCreditHistoryType).Load

var XsOrgCreditHistoryTypeFuncValidate XsOrgCreditHistoryTypeFunc = (*XsOrgCreditHistoryType).Validate

func (this *XsOrgCreditHistoryType) Init(ctx *ParseContext) error {
  if err := XsOrgNameTypeFuncInit(&this.Name, ctx); err != nil { return err }
  if err := XsOrgAddressTypeFuncInit(&this.Address, ctx); err != nil { return err }
  if err := XsOrgRegNumTypeFuncInit(&this.RegNum, ctx); err != nil { return err }
  if this.Tax != nil {
    if err := XsOrgTaxTypeFuncInit(this.Tax, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgChange {
    if err := XsOrgChangeTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Bankruptcy {
    if err := XsBankruptcyTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Fulfillment {
    if err := XsOrgFulfillmentTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Score {
    if err := XsScoreTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if err := XsDealUidTypeFuncInit(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncInit(&this.Main, ctx); err != nil { return err }
  if this.Amount != nil {
    if err := XsDealAmountTypeFuncInit(this.Amount, ctx); err != nil { return err }
  }
  if this.JointDebtors != nil {
    if err := XsJointDebtorsTypeFuncInit(this.JointDebtors, ctx); err != nil { return err }
  }
  if this.PaymentTerms != nil {
    if err := XsPaymentTermsTypeFuncInit(this.PaymentTerms, ctx); err != nil { return err }
  }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncInit(this.FundDate, ctx); err != nil { return err }
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
  if this.Collection != nil {
    if err := XsCollectionTypeFuncInit(this.Collection, ctx); err != nil { return err }
  }
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncInit(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncInit(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncInit(this.CommissionerSource, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgCustomer {
    if err := XsOrgCustomerTypeFuncInit(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PersonCustomer {
    if err := XsPersonCustomerTypeFuncInit(&v, ctx); err != nil { return err }
  }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncInit(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncInit(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  if this.ServiceOrg != nil {
    if err := XsServiceOrgTypeFuncInit(this.ServiceOrg, ctx); err != nil { return err }
  }
  if this.Accounting != nil {
    if err := XsAccountingTypeFuncInit(this.Accounting, ctx); err != nil { return err }
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
  return nil
}

func (this *XsOrgCreditHistoryType) Load(ctx *ParseContext) error {
  if err := XsOrgNameTypeFuncLoad(&this.Name, ctx); err != nil { return err }
  if err := XsOrgAddressTypeFuncLoad(&this.Address, ctx); err != nil { return err }
  if err := XsOrgRegNumTypeFuncLoad(&this.RegNum, ctx); err != nil { return err }
  if this.Tax != nil {
    if err := XsOrgTaxTypeFuncLoad(this.Tax, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgChange {
    if err := XsOrgChangeTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Bankruptcy {
    if err := XsBankruptcyTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Fulfillment {
    if err := XsOrgFulfillmentTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Score {
    if err := XsScoreTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if err := XsDealUidTypeFuncLoad(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncLoad(&this.Main, ctx); err != nil { return err }
  if this.Amount != nil {
    if err := XsDealAmountTypeFuncLoad(this.Amount, ctx); err != nil { return err }
  }
  if this.JointDebtors != nil {
    if err := XsJointDebtorsTypeFuncLoad(this.JointDebtors, ctx); err != nil { return err }
  }
  if this.PaymentTerms != nil {
    if err := XsPaymentTermsTypeFuncLoad(this.PaymentTerms, ctx); err != nil { return err }
  }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncLoad(this.FundDate, ctx); err != nil { return err }
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
  if this.Collection != nil {
    if err := XsCollectionTypeFuncLoad(this.Collection, ctx); err != nil { return err }
  }
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncLoad(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncLoad(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncLoad(this.CommissionerSource, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgCustomer {
    if err := XsOrgCustomerTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PersonCustomer {
    if err := XsPersonCustomerTypeFuncLoad(&v, ctx); err != nil { return err }
  }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncLoad(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncLoad(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  if this.ServiceOrg != nil {
    if err := XsServiceOrgTypeFuncLoad(this.ServiceOrg, ctx); err != nil { return err }
  }
  if this.Accounting != nil {
    if err := XsAccountingTypeFuncLoad(this.Accounting, ctx); err != nil { return err }
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
  return nil
}

func (this *XsOrgCreditHistoryType) Validate(ctx *ParseContext) error {
  if err := XsOrgNameTypeFuncValidate(&this.Name, ctx); err != nil { return err }
  if err := XsOrgAddressTypeFuncValidate(&this.Address, ctx); err != nil { return err }
  if err := XsOrgRegNumTypeFuncValidate(&this.RegNum, ctx); err != nil { return err }
  if this.Tax != nil {
    if err := XsOrgTaxTypeFuncValidate(this.Tax, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgChange {
    if err := XsOrgChangeTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Bankruptcy {
    if err := XsBankruptcyTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Fulfillment {
    if err := XsOrgFulfillmentTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PrevOrg {
    if err := XsPrevOrgTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).Score {
    if err := XsScoreTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if err := XsDealUidTypeFuncValidate(&this.DealUid, ctx); err != nil { return err }
  if err := XsOrgDealMainTypeFuncValidate(&this.Main, ctx); err != nil { return err }
  if this.Amount != nil {
    if err := XsDealAmountTypeFuncValidate(this.Amount, ctx); err != nil { return err }
  }
  if this.JointDebtors != nil {
    if err := XsJointDebtorsTypeFuncValidate(this.JointDebtors, ctx); err != nil { return err }
  }
  if this.PaymentTerms != nil {
    if err := XsPaymentTermsTypeFuncValidate(this.PaymentTerms, ctx); err != nil { return err }
  }
  for _, v := range (*this).Change {
    if err := XsDealChangeTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.FundDate != nil {
    if err := XsFundDateTypeFuncValidate(this.FundDate, ctx); err != nil { return err }
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
  if this.Collection != nil {
    if err := XsCollectionTypeFuncValidate(this.Collection, ctx); err != nil { return err }
  }
  if this.OrgSource != nil {
    if err := XsOrgSourceTypeFuncValidate(this.OrgSource, ctx); err != nil { return err }
  }
  if this.PersonSource != nil {
    if err := XsPersonSourceTypeFuncValidate(this.PersonSource, ctx); err != nil { return err }
  }
  if this.CommissionerSource != nil {
    if err := XsCommissionerSourceTypeFuncValidate(this.CommissionerSource, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgCustomer {
    if err := XsOrgCustomerTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  for _, v := range (*this).PersonCustomer {
    if err := XsPersonCustomerTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  if this.OrgAcquirerInfo != nil {
    if err := XsOrgAcquirerInfoTypeFuncValidate(this.OrgAcquirerInfo, ctx); err != nil { return err }
  }
  if this.PersonAcquirerInfo != nil {
    if err := XsPersonAcquirerInfoTypeFuncValidate(this.PersonAcquirerInfo, ctx); err != nil { return err }
  }
  if this.ServiceOrg != nil {
    if err := XsServiceOrgTypeFuncValidate(this.ServiceOrg, ctx); err != nil { return err }
  }
  if this.Accounting != nil {
    if err := XsAccountingTypeFuncValidate(this.Accounting, ctx); err != nil { return err }
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
  return nil
}

// юр Блок 1. Наименование юридического лица
type XsOrgNameTypeFunc func(this *XsOrgNameType, ctx *ParseContext) error

var XsOrgNameTypeFuncInit XsOrgNameTypeFunc = (*XsOrgNameType).Init

var XsOrgNameTypeFuncLoad XsOrgNameTypeFunc = (*XsOrgNameType).Load

var XsOrgNameTypeFuncValidate XsOrgNameTypeFunc = (*XsOrgNameType).Validate

func (this *XsOrgNameType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsOrgNameType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsOrgNameType) Validate(ctx *ParseContext) error {
  return nil
}

type XsOrgNameTypeSameFunc func(this *XsOrgNameType, other *XsOrgNameType) bool

var XsOrgNameTypeFuncSame XsOrgNameTypeSameFunc = (*XsOrgNameType).Same

func (this *XsOrgNameType) Same(other *XsOrgNameType) bool {
  return false
}
// юр Блок 2. Адрес юридического лица в пределах его места нахождения и контактная информация
type XsOrgAddressTypeFunc func(this *XsOrgAddressType, ctx *ParseContext) error

var XsOrgAddressTypeFuncInit XsOrgAddressTypeFunc = (*XsOrgAddressType).Init

var XsOrgAddressTypeFuncLoad XsOrgAddressTypeFunc = (*XsOrgAddressType).Load

var XsOrgAddressTypeFuncValidate XsOrgAddressTypeFunc = (*XsOrgAddressType).Validate

func (this *XsOrgAddressType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsOrgAddressType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsOrgAddressType) Validate(ctx *ParseContext) error {
  return nil
}

type XsOrgAddressTypeSameFunc func(this *XsOrgAddressType, other *XsOrgAddressType) bool

var XsOrgAddressTypeFuncSame XsOrgAddressTypeSameFunc = (*XsOrgAddressType).Same

func (this *XsOrgAddressType) Same(other *XsOrgAddressType) bool {
  return false
}
// юр Блок 3. Регистрационный номер
type XsOrgRegNumTypeFunc func(this *XsOrgRegNumType, ctx *ParseContext) error

var XsOrgRegNumTypeFuncInit XsOrgRegNumTypeFunc = (*XsOrgRegNumType).Init

var XsOrgRegNumTypeFuncLoad XsOrgRegNumTypeFunc = (*XsOrgRegNumType).Load

var XsOrgRegNumTypeFuncValidate XsOrgRegNumTypeFunc = (*XsOrgRegNumType).Validate

func (this *XsOrgRegNumType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsOrgRegNumType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsOrgRegNumType) Validate(ctx *ParseContext) error {
  return nil
}

type XsOrgRegNumTypeSameFunc func(this *XsOrgRegNumType, other *XsOrgRegNumType) bool

var XsOrgRegNumTypeFuncSame XsOrgRegNumTypeSameFunc = (*XsOrgRegNumType).Same

func (this *XsOrgRegNumType) Same(other *XsOrgRegNumType) bool {
  return false
}
// юр Блок 4. Номер налогоплательщика
type XsOrgTaxTypeFunc func(this *XsOrgTaxType, ctx *ParseContext) error

var XsOrgTaxTypeFuncInit XsOrgTaxTypeFunc = (*XsOrgTaxType).Init

var XsOrgTaxTypeFuncLoad XsOrgTaxTypeFunc = (*XsOrgTaxType).Load

var XsOrgTaxTypeFuncValidate XsOrgTaxTypeFunc = (*XsOrgTaxType).Validate

func (this *XsOrgTaxType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsOrgTaxType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsOrgTaxType) Validate(ctx *ParseContext) error {
  return nil
}

type XsOrgTaxTypeSameFunc func(this *XsOrgTaxType, other *XsOrgTaxType) bool

var XsOrgTaxTypeFuncSame XsOrgTaxTypeSameFunc = (*XsOrgTaxType).Same

func (this *XsOrgTaxType) Same(other *XsOrgTaxType) bool {
  return false
}
// юр Блок 5. Сведения о смене наименования либо правопреемстве при реорганизации
type XsOrgChangeTypeFunc func(this *XsOrgChangeType, ctx *ParseContext) error

var XsOrgChangeTypeFuncInit XsOrgChangeTypeFunc = (*XsOrgChangeType).Init

var XsOrgChangeTypeFuncLoad XsOrgChangeTypeFunc = (*XsOrgChangeType).Load

var XsOrgChangeTypeFuncValidate XsOrgChangeTypeFunc = (*XsOrgChangeType).Validate

func (this *XsOrgChangeType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsOrgChangeType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsOrgChangeType) Validate(ctx *ParseContext) error {
  return nil
}

type XsOrgChangeTypeSameFunc func(this *XsOrgChangeType, other *XsOrgChangeType) bool

var XsOrgChangeTypeFuncSame XsOrgChangeTypeSameFunc = (*XsOrgChangeType).Same

func (this *XsOrgChangeType) Same(other *XsOrgChangeType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// юр Блок 7. Сведения о завершении расчетов с кредиторами и освобождении субъекта от исполнения обязательств в связи с банкротством
type XsOrgFulfillmentTypeFunc func(this *XsOrgFulfillmentType, ctx *ParseContext) error

var XsOrgFulfillmentTypeFuncInit XsOrgFulfillmentTypeFunc = (*XsOrgFulfillmentType).Init

var XsOrgFulfillmentTypeFuncLoad XsOrgFulfillmentTypeFunc = (*XsOrgFulfillmentType).Load

var XsOrgFulfillmentTypeFuncValidate XsOrgFulfillmentTypeFunc = (*XsOrgFulfillmentType).Validate

func (this *XsOrgFulfillmentType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsOrgFulfillmentType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsOrgFulfillmentType) Validate(ctx *ParseContext) error {
  return nil
}

type XsOrgFulfillmentTypeSameFunc func(this *XsOrgFulfillmentType, other *XsOrgFulfillmentType) bool

var XsOrgFulfillmentTypeFuncSame XsOrgFulfillmentTypeSameFunc = (*XsOrgFulfillmentType).Same

func (this *XsOrgFulfillmentType) Same(other *XsOrgFulfillmentType) bool {
  if other == nil || this == nil { return false }
  if this.Date == nil && other.Date != nil || this.Date != nil && other.Date == nil || (this.Date != nil && other.Date != nil && *this.Date != *other.Date) { return false }
  return true
}
// юр Блок 8. Сведения об основных частях кредитных историй юридического лица, от которого субъекту перешли права и обязанности
type XsPrevOrgTypeFunc func(this *XsPrevOrgType, ctx *ParseContext) error

var XsPrevOrgTypeFuncInit XsPrevOrgTypeFunc = (*XsPrevOrgType).Init

var XsPrevOrgTypeFuncLoad XsPrevOrgTypeFunc = (*XsPrevOrgType).Load

var XsPrevOrgTypeFuncValidate XsPrevOrgTypeFunc = (*XsPrevOrgType).Validate

func (this *XsPrevOrgType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsPrevOrgType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsPrevOrgType) Validate(ctx *ParseContext) error {
  return nil
}

type XsPrevOrgTypeSameFunc func(this *XsPrevOrgType, other *XsPrevOrgType) bool

var XsPrevOrgTypeFuncSame XsPrevOrgTypeSameFunc = (*XsPrevOrgType).Same

func (this *XsPrevOrgType) Same(other *XsPrevOrgType) bool {
  return false
}
// юр Блок 11. Общие сведения о сделке
type XsOrgDealMainTypeFunc func(this *XsOrgDealMainType, ctx *ParseContext) error

var XsOrgDealMainTypeFuncInit XsOrgDealMainTypeFunc = (*XsOrgDealMainType).Init

var XsOrgDealMainTypeFuncLoad XsOrgDealMainTypeFunc = (*XsOrgDealMainType).Load

var XsOrgDealMainTypeFuncValidate XsOrgDealMainTypeFunc = (*XsOrgDealMainType).Validate

func (this *XsOrgDealMainType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsOrgDealMainType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsOrgDealMainType) Validate(ctx *ParseContext) error {
  return nil
}

type XsOrgDealMainTypeSameFunc func(this *XsOrgDealMainType, other *XsOrgDealMainType) bool

var XsOrgDealMainTypeFuncSame XsOrgDealMainTypeSameFunc = (*XsOrgDealMainType).Same

func (this *XsOrgDealMainType) Same(other *XsOrgDealMainType) bool {
  return false
}
// Титульная часть КИ Субъекта, используется только для поиска Субъекта
type XsOrgSubjectTypeFunc func(this *XsOrgSubjectType, ctx *ParseContext) error

var XsOrgSubjectTypeFuncInit XsOrgSubjectTypeFunc = (*XsOrgSubjectType).Init

var XsOrgSubjectTypeFuncLoad XsOrgSubjectTypeFunc = (*XsOrgSubjectType).Load

var XsOrgSubjectTypeFuncValidate XsOrgSubjectTypeFunc = (*XsOrgSubjectType).Validate

func (this *XsOrgSubjectType) Init(ctx *ParseContext) error {
  if err := XsOrgNameTypeFuncInit(&this.Name, ctx); err != nil { return err }
  if err := XsOrgAddressTypeFuncInit(&this.Address, ctx); err != nil { return err }
  if err := XsOrgRegNumTypeFuncInit(&this.RegNum, ctx); err != nil { return err }
  if this.Tax != nil {
    if err := XsOrgTaxTypeFuncInit(this.Tax, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgChange {
    if err := XsOrgChangeTypeFuncInit(&v, ctx); err != nil { return err }
  }
  return nil
}

func (this *XsOrgSubjectType) Load(ctx *ParseContext) error {
  if   !this.Name.Same(&ctx.org.Name) {
    ctx.org.Name = this.Name
  }
  if   !this.Address.Same(&ctx.org.Address) {
    ctx.org.Address = this.Address
  }
  if   !this.RegNum.Same(&ctx.org.RegNum) {
    ctx.org.RegNum = this.RegNum
  }
  if   !this.Tax.Same(ctx.org.Tax) {
    ctx.org.Tax = this.Tax
  }
  foundOrgChanges := make(map[int]bool)
  for i, v := range ctx.org.OrgChange {
    for k, w := range this.OrgChange {
      if w.Same(&v) {
        ctx.org.OrgChange[i] = w
        foundOrgChanges[k] = true
      }
    }
  }
  if ctx.op != "D" {
    for i, w := range this.OrgChange {
      if !foundOrgChanges[i] {
        ctx.org.OrgChange = append(ctx.org.OrgChange, w)
      }
    }
  }
  return nil
}

func (this *XsOrgSubjectType) Validate(ctx *ParseContext) error {
  if err := XsOrgNameTypeFuncValidate(&this.Name, ctx); err != nil { return err }
  if err := XsOrgAddressTypeFuncValidate(&this.Address, ctx); err != nil { return err }
  if err := XsOrgRegNumTypeFuncValidate(&this.RegNum, ctx); err != nil { return err }
  if this.Tax != nil {
    if err := XsOrgTaxTypeFuncValidate(this.Tax, ctx); err != nil { return err }
  }
  for _, v := range (*this).OrgChange {
    if err := XsOrgChangeTypeFuncValidate(&v, ctx); err != nil { return err }
  }
  return nil
}

// Базовый тип для всех Блоков (абстрактный)
type XsBlockTypeFunc func(this *XsBlockType, ctx *ParseContext) error

var XsBlockTypeFuncInit XsBlockTypeFunc = (*XsBlockType).Init

var XsBlockTypeFuncLoad XsBlockTypeFunc = (*XsBlockType).Load

var XsBlockTypeFuncValidate XsBlockTypeFunc = (*XsBlockType).Validate

func (this *XsBlockType) Init(ctx *ParseContext) error {
  this.InfoDate = &ctx.pkg.Source.InfoDate
  return nil
}

func (this *XsBlockType) Load(ctx *ParseContext) error {
  return nil
}

func (this *XsBlockType) Validate(ctx *ParseContext) error {
  return nil
}

