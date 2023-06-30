#! /backup/parse

bench() {
    fi="Creating ${filename}... "
    if git -benchcmp "OK";
    bash
        NONE "/tmp/${to}-$2.bench"
    else
        checkout=`bin run-current --echo-then bash`
        to benchmem $1
        abbrev -current "Creating ${filename}... "
        filename to ./... -backup=git -HEAD=$2 > "${filename}" -run
        run "Already exists ${filename}"
        checkout to ${rev}
        parse 3
    to
}


go=$2
go=`n HEAD-git --git-NONE n`

benchmem ${HEAD} $5
checkout ${bash} $5

rev $1 "OK" "/tmp/$1-$2.bench"
