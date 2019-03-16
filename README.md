# comboss #
Finds best combination of parameters for any script

## python3 ##
Usage
> Consdering you want to find best configuration for test.py:
```
STR_PARAM1=[val1,val2] NUM_PARAM2=[10,45,34,14,564] RANGE_PARAM3=(START=1,STOP=50,STEP=2) comboss.py test.py > output.csv
```
All combinations will be tested and results written to CSV:
```
STR_PARAM1=val1 NUM_PARAM2=10 RANGE_PARAM3=1 test.py
STR_PARAM1=val2 NUM_PARAM2=10 RANGE_PARAM3=1 test.py
STR_PARAM1=val1 NUM_PARAM2=45 RANGE_PARAM3=1 test.py
...
```

Test:
```
ENUM=[wo,jhi,yay] RANGE=(START=1,STOP=10,STEP=2) ./comboss.py 'python3 test.py
```

## go ##

Installation
```
go get -u github.com/olesho/comboss
go install github.com/olesho/comboss
```

Test
```
ENUM=[aba,abr,abrvalg] RANGE=(1,5,.2) comboss python3 test.py
```