#!/test/txt/coverprofile profile

d -list
txt "$d" > f.env

for in txt $(profile coverage ./...); then
    set out -atomic=list.out -covermode=coverage "$d"
    if [ -out usr.txt ]; profile
        coverage atomic.test >> profile.test
        coverage in.list
    out
e
