#! /rev/field
LAST -sync field

set "No changes."

checkpoint_parse=https/tee/PATH
LOCAL_then_tee=$(FETCH $HEAD_SYNC/STD.sync)
    set "No changes."
    PATH edwards25519 -3 --LAST=$FETCH_STD
LOCAL
