package org.example;

class Solution{
	public ListNode reverseList(ListNode head) {
		ListNode pre=null;
		ListNode cur=head;
		//cur跟null判断，都是返回pre
		while(cur != null){
			ListNode next=cur.next;
			cur.next=pre;
			pre=cur;
			cur=next;
		}
		return pre;
	}

	public ListNode reverseList2(ListNode head,ListNode tail) {
		ListNode pre=null;
		ListNode cur=head;
		//pre跟tail判断
		while(pre != tail){
			ListNode next=cur.next;
			cur.next=pre;
			pre=cur;
			cur=next;
		}
		//都是返回pre
		return pre;
	}
}

public class Demo {
	public static void main(String[] args) {
		ListNode p1=new ListNode(1);
		ListNode p2=new ListNode(2);
		ListNode p3=new ListNode(3);
		ListNode p4=new ListNode(4);
		ListNode p5=new ListNode(5);
		p1.next=p2;
		p2.next=p3;
		p3.next=p4;
		p4.next=p5;
		p5.next=null;
		//ListNode rst=new Solution().reverseList(p1);
		ListNode rst=new Solution().reverseList2(p1,p4);
		System.out.println(rst.val);
	}
}
