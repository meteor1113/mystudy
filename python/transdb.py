import sqlite3
import zlib

conn=sqlite3.connect(r"f:\input.db")
c=conn.cursor()

conn2=sqlite3.connect(r"f:\input2.db")
c2=conn2.cursor()

for row in c.execute("select * from TestTree"):
    print row[0]
    #print type(row[1])
    content = row[1].encode('utf-8')
    s=zlib.compress("blob %d\0" % len(content) + content)
    
    #print s
    #sql = "update TestTree set Content='%s' where Key='%s'" % (sqlite3.Binary(s), row[0])
    #print sql
    print "c2.execute()"
    c2.execute("insert into TestTree values(?, ?)", (row[0], buffer(s)))

c.close()
conn.close()

conn2.commit()
c2.close()
conn2.close()

    
