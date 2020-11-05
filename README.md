# TABLE TO STRUCT PARSER

## OVERVIEW

Golang does not include any ORM repositories native to the language. However, there exists multiple 3rd party packages which can be used as an object relationship mapper for different databases. This application was created in order to read from a table, extract the column names and datatypes and create a .go file which can be used.

This was created in order to speed up the process of creating an exceptionally large struct to represent the table of an already existing table.