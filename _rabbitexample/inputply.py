filename = 'bunny.ply'

points = []
tris = []
with open(filename, 'r') as f:
	lines = f.readlines()
	pnum = int(lines[4].split()[2])
	print(pnum)
	fnum = int(lines[10].split()[2])
	print(fnum)
	pbase = 13
	for i in range(pnum):
		points.append(tuple(map(float,lines[pbase+i].split()))[:3])
	fbase = pbase+pnum
	for i in range(fnum):
		c,x,y,z = map(int,lines[fbase+i].split())
		assert(c==3)
		tris.append((points[x],points[y],points[z]))
	s = ''
	for x in tris:
		s+= '%f %f %f %f %f %f %f %f %f\n'%(*x[0],*x[1],*x[2])
	with open('tmp.txt', 'w') as out:
		out.write(s)

