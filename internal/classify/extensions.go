package classify

// folders is just a slice of strings containing the names of the subdirectories
// to be created in the target directory.
var folders = []string{"Images", "Music", "Videos", "Documents", "Archives"}

// extensions is a map linking the extensions to their respective subdirectories.
//
// I used this approach so looking up the correct directory is an O(1) operation
// without a messy switch-case. Although this would probably prove to look verbose
// but I think I am okay with this. Plus adding an extension is simply just an extra
// line here
var extensions = map[string]string{
	".jpg":  "Images",
	".jpeg": "Images",
	".png":  "Images",
	".gif":  "Images",
	".webp": "Images",
	".tiff": "Images",
	".tif":  "Images",
	".ico":  "Images",
	".heic": "Images",
	".svg":  "Images",
	".mp3":  "Music",
	".m4a":  "Music",
	".wav":  "Music",
	".flac": "Music",
	".aac":  "Music",
	".opus": "Music",
	".ogg":  "Music",
	".mp4":  "Videos",
	".mkv":  "Videos",
	".avi":  "Videos",
	".mov":  "Videos",
	".webm": "Videos",
	".pdf":  "Documents",
	".doc":  "Documents",
	".docx": "Documents",
	".xls":  "Documents",
	".xlsx": "Documents",
	".ppt":  "Documents",
	".pptx": "Documents",
	".txt":  "Documents",
	".csv":  "Documents",
	".md":   "Documents",
	".epub": "Documents",
	".zip":  "Archives",
	".tar":  "Archives",
	".gz":   "Archives",
	".7z":   "Archives",
	".rar":  "Archives",
	".tgz":  "Archives",
}
