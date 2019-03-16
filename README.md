# tuner #
Finds best combination of parameters for any script

Usage
> Consdering you want to find best configuration for test.py:
```
STR_PARAM1=[val1,val2] NUM_PARAM2=[10,45,34,14,564] RANGE_PARAM3=(START=1,STOP=50,STEP=2) tuner.py test.py > output.csv
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
ENUM=[wo,jhi,yay] RANGE=(START=1,STOP=10,STEP=2) ./tuner.py 'python3 test.py
```