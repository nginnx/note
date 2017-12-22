#!/bin/awk
awk 'BEGIN {temp="";int1 = 0};
{
if($0~/^commit/){
	if(int1==1){
	print temp "\n";
	}
	temp = ""
	int1 = 1;
}

if($0~/^Author/){
	if($0~/.*xukq.*/){
	int1 =1;
	}else{
	temp=""
	int1 =0;
	}
}
if(int1 ==1){
	temp=temp $0 "\n";
}

}; END{print temp;}'
