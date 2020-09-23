# Collection.


## Short Description.

This Package provides a Collection Functionality.

## Full Description.

The Collection is stored in the SQL Database.
The Collection consists of Classes. 
Each Class has Objects and Reference Properties (Property Types).
Each Class Object has Class Object Properties with Values and a Link to a Reference Class Property.
Each Reference Class Property has its Type (Kind).

The Package provides Functionality to store and read wide Collections. 
Saying 'wide' here means that the Properties Set may be very flexible.
Each Property stores its Values in a separate Database Table.
The above stated Feature allows to economize space for non set Properties and 
to decrease Access Time for simple Queries when only a single Property is requested.
Also, such an Approach increases the reliability of the Storage. 
If something bad happens with one Property, other Properties will not be touched.

This Library is a simple Example of what can be done. The Functionality is not full here.

## Installation.

Import Commands:
```
go get -u "github.com/legacy-vault/library/go/collection/Uc"
```

## Usage.

Usage Example can be found in the Tests:

[https://github.com/legacy-vault/library/tree/master/go/collection/UcTest](https://github.com/legacy-vault/library/tree/master/go/collection/UcTest)
