compile-luas:
	for f in tests/*.lua; do luac5.1 -o "$${f}.out" "$${f}"; done