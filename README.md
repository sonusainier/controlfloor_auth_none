# controlfloor_auth_none
Passthru auth for controlfloor

How to use:
* Clone controlfloor repo
* Clone this repo
* Alter controlfloor/go.mod
* Add line `replace github.com/nanoscopic/controlfloor_auth => /your/path/to/controlfloor_auth_none` above the `require` block.