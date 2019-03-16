#!/usr/bin/env python

import os
import re
import sys
import subprocess

def isfloat(value):
	try:
		float(value)
		return True
	except ValueError:
		return False

def isint(s):
	try: 
		int(s)
		return True
	except ValueError:
		return False

params = {}
for k,v in os.environ.items():
	if v.startswith('[') and v.endswith(']'):
		parts = v[1:-1].split(',')
		params[k] = parts
	if v.startswith('(') and v.endswith(')'):
		parts = v[1:-1].split(',')
		if len(parts) == 3:
			start = parts[0]
			stop = parts[1]
			step = parts[2]
			if len(start) > 6:
				if start[:6].lower() == 'start=':
					start = start[6:]
			if len(stop) > 5:
				if stop[:5].lower() == 'stop=':
					stop = stop[5:]
			if len(step) > 5:
				if step[:5].lower() == 'step=':
					step = step[5:]

			if all([isint(start), isint(stop), isint(step)]):
				start = int(start)
				stop = int(stop)
				step = int(step)
			elif all([isfloat(start), isfloat(stop), isfloat(step)]):
				start = float(start)
				stop = float(stop)
				step = float(step)
		sets = []
		n = start
		while n <= stop:
			sets += [str(n)]
			n += step
		params[k] = sets

keys = [k for k, _ in params.items()]

def next_param(combos, param_index=0, current_combo={}):
	if param_index == len(params):
		combos += [current_combo]
		return
	for val in list(params[keys[param_index]]):
		current_combo[keys[param_index]] = val
		next_param(combos, param_index+1, dict(current_combo)) 

combos = []
next_param(combos)

print(combos)

processes = []
for combo in combos:
		processes.append(subprocess.Popen(sys.argv[1].split(), env=dict(os.environ, **combo)))
for p in processes:
		p.wait()
