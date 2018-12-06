input_file = ".\bestbuy-buybox.xlsx"
output_file = ".\bestbuy-buybox.txt"
tsv_format = -4158

Set objFSO = CreateObject("Scripting.FileSystemObject")

src_file = objFSO.GetAbsolutePathName(input_file)
dest_file = objFSO.GetAbsolutePathName(output_file)

Dim oExcel
Set oExcel = CreateObject("Excel.Application")

Dim oBook
Set oBook = oExcel.Workbooks.Open(src_file)

oBook.SaveAs dest_file, tsv_format

oBook.Close False
oExcel.Quit