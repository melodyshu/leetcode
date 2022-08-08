class Solution {
    public String minWindow(String s, String t) {
        Map<Character,Integer> need = new HashMap<Character,Integer>();
        Map<Character,Integer> window = new HashMap<Character,Integer>();
        int left = 0,right = 0;
        int start = 0;
        int count = 0;
        int len = Integer.MAX_VALUE;
        //初始化need
        for(Character c : t.toCharArray()){
            need.put(c,need.getOrDefault(c,0)+1);
        }
        //扩大窗口
        while(right < s.length()){
            //更新window
            char c = s.charAt(right);
            right++;
            if(need.containsKey(c)){
                window.put(c,window.getOrDefault(c,0)+1);
                if(need.get(c).equals(window.get(c))){
                    count++;
                    }
            }

            //缩小窗口
            while(count == need.size()){
                if(right - left < len){
                    len=right - left;
                    start = left;
                }
                char d = s.charAt(left);
                left++;
                if(need.containsKey(d)){
                    if(need.get(d).equals(window.get(d))){
                        count--;
                    }
                    window.put(d,window.get(d)-1);
                }
                
            }
    }
    return len == Integer.MAX_VALUE?"":s.substring(start,start+len);
    }
}