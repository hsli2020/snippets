if WScript.Arguments.Count < 3 Then
    WScript.Echo "Please specify the sheet, the source, the destination files. Usage: ExcelToCsv <sheetName> <xls/xlsx source file> <csv destination file>"
    Wscript.Quit
End If

csv_format = 6

Set objFSO = CreateObject("Scripting.FileSystemObject")

src_file = objFSO.GetAbsolutePathName(Wscript.Arguments.Item(1))
dest_file = objFSO.GetAbsolutePathName(WScript.Arguments.Item(2))

Dim oExcel
Set oExcel = CreateObject("Excel.Application")

Dim oBook
Set oBook = oExcel.Workbooks.Open(src_file)

oBook.Sheets(WScript.Arguments.Item(0)).Select
oBook.SaveAs dest_file, csv_format

oBook.Close False
oExcel.Quit