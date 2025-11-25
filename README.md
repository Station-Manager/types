# Station Manager: types package

This package contains the types used across the various modules and applications which comprise *Station Manager*
application suite.

It is a core module that defines the data structures and types that are shared and
**does not/cannot depend on any other **Station Manager package**; only Go system modules (e.g. "time") are allowed.
It is essential that this module **HAS NO DEPENDENCIES** from *Station Manager* to ensure that it can be used in
anywhere without causing cyclic dependency errors.

Candidates for inclusion in this module are:
1. Data structures and types that are shared across modules/packages
2. Enumerations that are used across modules/packages
3. Constants that are used across modules/packages
4. Interfaces that are implemented by multiple modules/packages

