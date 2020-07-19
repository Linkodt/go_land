package main

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:"红孩儿",
		Age:10,
		Skill:"喷火",
	}
	res := monster.Store()
	if !res{
		t.Fatalf("monster.Store() fail!!, want=%v  real=%v", true, res)
	}
	t.Logf("monster.Store() success!!!")
}

func TestRestore(t *testing.T){
	var monster = &Monster{}
	res := monster.Restore()
	if !res{
		t.Fatalf("monster.Restore() fail!!, want=%v  real=%v", true, res)
	}
	if monster.Name != "红孩儿"{
		t.Fatalf("monster.Restore() fail!!, want=%v  real=%v", "红孩儿", monster.Name)
	}
	t.Log(*monster)
	t.Logf("monster.Restore() success!!!")
}