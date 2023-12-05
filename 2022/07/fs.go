package main


type FSChild interface{
	GetSize() uint64
	IsDir() bool
	GetParent() *Dir
	GetName() string
}

/*
DIRECTORY
*/

type File struct{
	Name string
	Parent *Dir
	Size uint64
}

func NewFile(name string, size uint64, parent *Dir) *File{
	return &File{
		Name: name,
		Size: size,
		Parent:parent,
	}
}

func (f *File) GetSize() uint64 {
	return f.Size
}

func (f *File) IsDir() bool {
	return false
}

func (f *File) GetParent() *Dir {
	return f.Parent	
}

func (f *File) GetName() string {
	return f.Name
}

/*
DIRECTORY
*/
type Dir struct{
	Name string
	Size uint64
	Files []*File
	Directories []*Dir
	Parent *Dir
}

func NewDir(name string, parent *Dir) *Dir {
	dir := Dir{
		Name: name,
		Parent: parent,
	}
	return &dir
}

func (d *Dir) GetSize() uint64 {
	return d.Size
}

func (d *Dir) IsDir() bool {
	return true
}

func (d *Dir) GetParent() *Dir {
	return d.Parent	
}

func (d *Dir) GetName() string {
	return d.Name
}

// func (d *Dir) ProcessChildren() uint64 {
//
// }

func (d *Dir) MoveOut() *Dir {
	return d.GetParent()
}

func (d *Dir) GetDirectories() []*Dir {
	return d.Directories
}
func (d *Dir) GetFiles() []*File {
	return d.Files
}

func (d *Dir) AddDir(dir *Dir) {
	d.Directories = append(d.Directories, dir)
}
func (d *Dir) AddFile(file *File){
	d.Files = append(d.Files, file)
}

func (d *Dir) IncSize(childsSize uint64) uint64 {
	d.Size += childsSize
	return d.Size
}

func (d *Dir) IncSizeFromSubDirs(){
	for _,dir := range d.GetDirectories() {
		if len(dir.GetDirectories()) > 0 {
			dir.IncSizeFromSubDirs()	
		}
		d.IncSize(dir.GetSize())
	}
}

func (d *Dir) ListChildrenTotals() []uint64 {
	
}


func (d *Dir) FindDir(dirname string) *Dir{
	for _, dir := range d.GetDirectories(){
		if dir.GetName() == dirname {
			return dir
		}
	}
	panic("Couldnt find SubDirectory")
}

func getname(fsc FSChild) string{
	return fsc.GetName()
}

