# Fixed Size Bubble Cache.


## Short Description.

This Package provides a fixed Size Bubble Cache Functionality.

## Full Description.

Stores Information about last N active Records,
where 'N' is a fixed Number of Records.

The Cache Object has two Parameters:

	*	Maximum Size (N, mentioned above);
	*	TTL (Period is set in Seconds).

When we add a Record to the Cache, if an incoming Record already exists in the 
Cache, it is moved from its existing Position to the Top of the Cache. The Term 
'exists' means that there is a Record in the Cache with the same UID as the UID 
of the inserted Record.

Each Record has a 'UID' and a 'Data' Fields.
'UID' is used for Indexing. 'Data' is used to store some useful Information.
It is not recommended to set the 'Data' Field as 'nil'. If you do not want to 
store anything, it is better to set it to 'true' of 'false'.

If an incoming Record is new (does not exist in the Cache), it is added to the 
Top of the Cache deleting the Bottom Element from the Cache (all the existing 
Cache Items are virtually "moved" 1 Step down in the "Ladder" of Records). Of 
course, all the "moved" Records are not moved anywhere, while this Cache uses a 
double Link List to store Records, which makes Insertion and Removal 
Operations very fast.

Example:
[ghi] + [abc,def,ghi,jkl,xyz] => [ghi,abc,def,jkl,xyz].
[xxx] + [abc,def,ghi,jkl,xyz] => [xxx,abc,def,ghi,jkl].

When the User requests a Value (by the UID) from the Cache, we first, check its 
Existence in the Cache's List, and second, check Record's TTL (Time To Live). 
If the requested Record exists but is outdated, we remove it from the Cache.

The Removals are done in a "Lazy" Style: either when the Record is requested, or
when a new Record arrives and we have no free Space to store old Records. This 
is done to avoid Usage of Workers inspecting big amounts of Data in Real-Time 
and to save a lot of CPU Time. We check TTL only when it is neccessary.

## Installation.

Import Commands:
```
go get "github.com/legacy-vault/library/go/fixed_size_bubble_cache"
go get "github.com/legacy-vault/example/go/fixed_size_bubble_cache/code"
```

## Usage.

Usage Example can be found at the following Address:

[https://github.com/legacy-vault/example/tree/master/go/fixed_size_bubble_cache](https://github.com/legacy-vault/example/tree/master/go/fixed_size_bubble_cache)
