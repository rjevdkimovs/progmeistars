package main
import ("fmt")
func main(){
	s:="Rīgas torņī";
	for(_, r:=range[]rune(s),_){
		fmt.Printf("%c %v \n",EndcodeRune(r))
	}
	var r rune
	r=0x110000
	fmt.Printf("\n valid rune:%c %v \n", r, EndcodeRune(r))
}
func EndcodeRune(r rune)[] byte {
	if r<0x80{
		return[]byte{byte(r)}
	} else
	if r<0x800{
		return[]byte{
			0xc0|byte(r>>6),
			0x80|byte(r&0x3f)}
		} else
		if r<0x10000{
			return[]byte{
				0xe0|byte(r>>12),
				0x80|byte(r>>6&0x3f),
				0x80|byte(r&0x3f)}
			} else
			if r<0x110000{
				return[] byte{
					0xf0|byte(r>>18),
					0x80|byte(r>>12 & 0x3f),
					0x80|byte(r>>6 & 0x3f),
					0x80|byte(r>>6 & 0x3f)}
				} else {
					return[] byte{}
					
	
}
