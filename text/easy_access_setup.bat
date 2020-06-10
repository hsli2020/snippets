mklink /d staff \\bteserver16-pc\f\staff
mklink /d returnlabel "\\bteserver16-pc\f\All return label to customers"
mklink /d borderworx "\\bteserver16-pc\f\BorderworxUS-Shipments"
mklink /d bteshipping "\\bteserver16-pc\f\BTE-Shipping"
mklink /d damagedproductphoto "\\bteserver16-pc\f\Damaged product photots"
mklink /d HP_M525_Scan "\\bteserver16-pc\f\HP_M525_Scan"
mklink /d FBA_compliants "\\bteserver16-pc\f\FBA_compalints"
mklink /d damaged_product_photo "\\bteserver16-pc\f\Damaged product photos"
mklink /d marketing "\\bteserver16-pc\f\Marketing"
mklink /d ups_return_label "\\bteserver16-pc\f\ups-return-label"
mklink /d bte_pricelist "\\bteserver16-pc\g\\BTE-Price-list"
mklink /d fba_planning "\\bteserver16-pc\g\\FBA planning"
mklink /d purchasing "\\bteserver16-pc\g\\Purchasing"
mklink /d sbn_planning "\\bteserver16-pc\g\\SBN Planning"
mklink /d dataexchange_output "\\bteserver16-pc\DataExchange\out"

fsutil behavior set SymlinkEvaluation L2R:1
fsutil behavior set SymlinkEvaluation R2R:1
fsutil behavior set SymlinkEvaluation R2L:1
fsutil behavior set SymlinkEvaluation L2L:1