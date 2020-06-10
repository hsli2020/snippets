@echo off

set hour=%time:~0,2%

rem EQU - equal
rem NEQ - not equal
rem LSS - less than
rem LEQ - less than or equal
rem GTR - greater than
rem GEQ - greater than or equal

if %hour% GEQ 16 if %hour% LEQ 18 (
    echo php InboundToStockJob.php
)

if %hour% GEQ 10 if %hour% LEQ 20 (
    echo php WorldshipAddressBookJob.php
    echo php FedexAddressBookJob.php
    echo php CanadaPostAddressBookJob.php
    echo php DHLAddressBookJob.php
)
