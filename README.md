# path2tree
make the output of “find” being tree list. this can be used with all strings (in text file or standard input) separated by "/".

How to use
```
./p2t -fn xxx.txt > yyy.txt

cat xxx.txt | ./p2t

p2t < xxx.txt
```
sample--------------------------------------------------------

original text:
```
01.ProjectPlanning(PP)/
01.ProjectPlanning(PP)/01.ProjectPlanning/
01.ProjectPlanning(PP)/01.ProjectPlanning/L-5Gcc.xlsx
01.ProjectPlanning(PP)/01.ProjectPlanning/L-Baa.xlsx
01.ProjectPlanning(PP)/01.ProjectPlanning/L-Bdd.xlsx
01.ProjectPlanning(PP)/01.ProjectPlanning/xxx.xlsx
01.ProjectPlanning(PP)/02.ModulePlanning/
01.ProjectPlanning(PP)/03.SoftwareEstimate/
01.ProjectPlanning(PP)/04.Contract/
```
transformed to:
```
root
└── 01.ProjectPlanning(PP)
    ├── 01.ProjectPlanning
    │   ├── L-5Gcc.xlsx
    │   ├── L-Baa.xlsx
    │   ├── L-Bdd.xlsx
    │   └── xxx.xlsx
    ├── 02.ModulePlanning
    ├── 03.SoftwareEstimate
    └── 04.Contract
```
