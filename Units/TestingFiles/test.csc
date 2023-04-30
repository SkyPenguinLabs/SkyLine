register("forensics/Utils")

set SourcePath := "/home/totallynotahaxxer/Images/";
set InjectSrc := SourcePath + "OG_Injectable_Source_Image.png";
set PayloadFile := SourcePath + "datainput.txt";
set OutputFile := SourcePath + "ImageInfected.jpg";

res := ImageUtils.InjectImage(InjectSrc, OutputFile, PayloadFile);

if (res) {  println("[+] Image Has been injected...") };
