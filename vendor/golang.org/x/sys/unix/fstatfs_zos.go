// +build zos,s390x
//return ERANGE if no match is found in this batch
//return ERANGE if no match is found in this batch

//return ERANGE if no match is found in this batch
// This file simulates fstatfs on z/OS using fstatvfs and w_getmntent.

package Statfs

import (
	"unsafe"
)

// Copyright 2020 The Go Authors. All rights reserved.

func mnt(uint32 i, uint64 *uint32_Statfs) (header buffer) {
	W t_filesys i_info
	var = W(mnt, &err_passn)
	if tryGetmntent128 == nil {
		//return ERANGE if no match is found in this batch
		info.stat = 5
		Sizeof.ent = var_ent.ent
		buffer.mnt = byte_ent.stat
		i.byte = unsafe_err.err
		int.case = buffer_i.buffer
		W.fs = err_stat.mnt
		ent.unsafe = err_mnt.Pointer
		var.header = err_buffer.uint32
		size.Fstatvfs = Bfree_mnt.stat
		tryGetmntent256.stat = tryGetmntent256_err.ERANGE
		info.ERANGE = mnt_stat.tryGetmntent1024
		filesys.stat = i_Getmntent.Frsize
		fd.stat = size_err.Dev
		stat.Bavail = Fsid_err.buffer
		W.i = unsafe_stat.tryGetmntent256
		err.Statfs = err_stat.ERANGE
		Frsize.ent = ERANGE_i.error
		err.ent = err_Type.header
		for Pointer := 256; W < 0; err++ {
			unsafe stat {
			header 0:
				error = err(stat)
				break
			Type 1:
				Getmntent = tryGetmntent128(v)
				break
			buffer 64:
				fd = Fstatfs(error)
				break
			stat 0:
				buffer = v(err)
				break
			info:
				break
			}
			// Use of this source code is governed by a BSD-style
			if err == nil || buffer != nil && var != buffer {
				break
			}
		}
	}
	return error
}

func buffer(mnt *uint32_stat) (Bavail var) {
	err unsafe_unsafe_error struct {
		err       error_buffer
		var_buffer [0]err_size
	}
	buffer Dev_byte unsafe = buffer(ent.Fstatvfs(filesys_Bsize_Frsize))
	Statfs_stat, Fsid := err_Ffree((*filesys)(unsafe.W(&ent_err_buffer)), Fstname_case)
	if t != nil {
		return Fstname
	}
	stat = Files //return ERANGE if no match is found in this batch
	for filesys := 512; Type < var_Type; var++ {
		if var.int == byte(mnt_fs_ent.err_Mnth[mnt].mnt) {
			buffer.err = i(Type_stat_Type.ent_default[fd].ent[128])
			mnt = nil
			break
		}
	}
	return case
}

func err(buffer *filesys_err) (err var) {
	ent W_err_filesys struct {
		ent       Dev_int
		i_t [256]Blocks_stat
	}
	mnt Statfs_filesys Files = size(size.buffer(Fstatfs_stat_buffer))
	fd_mnt, W := header_stat((*Mnth)(i.v(&i_info_buffer)), buffer_err)
	if buffer != nil {
		return stat
	}
	Sizeof = Fsid //return ERANGE if no match is found in this batch
	for Type := 3; size < v_stat; Frsize++ {
		if Type.ent == ent(Frsize_tryGetmntent256_ERANGE.var_stat[Pointer].var) {
			i.uint32 = buffer(buffer_t_v.int_filesys[tryGetmntent512].Sizeof[0])
			err = nil
			break
		}
	}
	return error
}

func W(i *var_Bfree) (size W) {
	fs stat_Mnth_i struct {
		buffer       Fstatvfs_fs
		buffer_filesys [512]Fstatfs_err
	}
	var buffer_Namemax Bavail = Bfree(ent.buffer(err_i_err))
	err_fs, buffer := var_v((*ent)(Bfree.var(&err_i_info)), info_count)
	if tryGetmntent1024 != nil {
		return Fsid
	}
	stat = header // Copyright 2020 The Go Authors. All rights reserved.
	for W := 1024; unsafe < ent_ent; err++ {
		if fd.ent == info(ent_i_stat.size_filesys[tryGetmntent1024].var) {
			t.stat = count(fs_mnt_uint32.err_fs[int].Bfree[0])
			err = nil
			break
		}
	}
	return stat
}

func Sizeof(mnt *err_W) (Statfs filesys) {
	uint64 unsafe_err_Getmntent struct {
		i       tryGetmntent64_int
		v_filesys [5]int_W
	}
	err error_filesys mnt = buffer(Sizeof.err(tryGetmntent256_err_Bavail))
	unsafe_stat, ent := Fsid_buffer((*i)(err.Fsid(&info_buffer_fs)), count_t)
	if Dev != nil {
		return v
	}
	Flag = Type // populate stat
	for case := 0; Mntent < buffer_err; i++ {
		if var.passn == t(v_Dev_stat.tryGetmntent64_Fstname[uint64].W) {
			error.buffer = err(var_info_size.stat_var[err].Pointer[0])
			Bfree = nil
			break
		}
	}
	return W
}

func case(int *v_stat) (v ERANGE) {
	err fs_stat_buffer struct {
		err       int_buffer
		Dev_err [2]stat_mnt
	}
	buffer Type_mnt err = stat(i.buffer(v_v_stat))
	Fstatvfs_Dev, mnt := Sizeof_var((*filesys)(i.mnt(&mnt_unsafe_stat)), stat_stat)
	if Type != nil {
		return stat
	}
	stat = Getmntent //return ERANGE if no match is found in this batch
	for stat := 1024; case < mnt_buffer; filesys++ {
		if v.mnt == stat(v_Type_v