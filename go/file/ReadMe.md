# File and Folder Listing.


## Short Description.

This Package provides File and Folder Listing Functionality.

## Full Description.

This Package provides File and Folder Listing Functionality.<br />
It can list only Files, only Folder and both of them together in various Modes.<br />
Available Modes are the following:<br />
  - Contents of current Directory;
  - Contents of current Directory and all Sub-Directories;
  - Contents of the Directory with Hierarchy Level Number equal to N;
  - Contents of the Directory with Hierarchy Level Number less than or equal to N;
  - Contents of the Directory with Hierarchy Level Number greater than or equal to N.

<br />
Directory with Hierarchy Level Number is counted from the current Directory, where the current Directory has a virtual Number Zero (it can not be really used in Practice), and all its Contents having a positive Level Number.<br />
<br />
Also this Package can list Files from a Directory (with or without its Sub-Directories) applying simple Restrictions to the File's Extension. Restrictions may be of two Types: allowed Extensions and forbidden Extensions.<br />
<br />
As a Bonus, this Package provides a Collection of exotic Functions to list Edge Files and Edge Folders. Here, 'Edge' means that it is the "End-of-the-Line" Item, which has no further Sub-Levels. These Functions may be helpful in some exotic Cases for storing huge Amounts of Data using the File System Folders as Database Indices. For Example, some old Web Hosting Services used to store User Data in such Folders as '/storage/type/subtype/genre/style/.../content_folder', where Folder Names were used as Indices in a long virtual Address.

## Installation.

Import Commands:
```
go get -u "github.com/legacy-vault/library/go/file"
go get -u "github.com/legacy-vault/example/go/file/code"
```

## Usage.

Usage Example can be found at the following Address:

[https://github.com/legacy-vault/example/tree/master/go/file](https://github.com/legacy-vault/example/tree/master/go/file)
