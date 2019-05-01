package main

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	type test struct {
		name string
		args args
		want *ListNode
	}
	tests := []test{
		func() test {
			tail := &ListNode{
				Val: -4,
			}
			want := &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  0,
					Next: tail,
				},
			}
			tail.Next = want
			return test{
				name: "head = [3,2,0,-4], pos = 1",
				args: args{
					head: &ListNode{
						Val:  3,
						Next: want,
					},
				},
				want: want,
			}
		}(),
		func() test {
			tail := &ListNode{
				Val: 2,
			}
			want := &ListNode{
				Val:  1,
				Next: tail,
			}
			tail.Next = want
			return test{
				name: "head = [1, 2], pos = 0",
				args: args{
					head: want,
				},
				want: want,
			}
		}(),
		func() test {
			return test{
				name: "head = [1, 2], pos = 0",
				args: args{
					head: &ListNode{
						Val: 1,
					},
				},
				want: nil,
			}
		}(),
		func() test {
			tail := &ListNode{
				Val: -5,
			}
			want := &ListNode{
				Val: -9,
				Next: &ListNode{
					Val: -5,
					Next: &ListNode{
						Val:  -2,
						Next: tail,
					},
				},
			}
			tail.Next = want
			return test{
				name: "head = [-1,-7,7,-4,19,6,-9,-5,-2,-5], pos = 6",
				args: args{
					head: &ListNode{
						Val: -1,
						Next: &ListNode{
							Val: -7,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: -4,
									Next: &ListNode{
										Val: 19,
										Next: &ListNode{
											Val:  6,
											Next: want,
										},
									},
								},
							},
						},
					},
				},
				want: want,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
