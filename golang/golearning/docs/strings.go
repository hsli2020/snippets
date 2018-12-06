Package strings
===============
    func Compare(a, b string) int
    func Contains(s, substr string) bool
    func ContainsAny(s, chars string) bool
    func ContainsRune(s string, r rune) bool
    func Count(s, sep string) int
    func EqualFold(s, t string) bool
    func Fields(s string) []string
    func FieldsFunc(s string, f func(rune) bool) []string
    func HasPrefix(s, prefix string) bool
    func HasSuffix(s, suffix string) bool
    func Index(s, sep string) int
    func IndexAny(s, chars string) int
    func IndexByte(s string, c byte) int
    func IndexFunc(s string, f func(rune) bool) int
    func IndexRune(s string, r rune) int
    func Join(a []string, sep string) string
    func LastIndex(s, sep string) int
    func LastIndexAny(s, chars string) int
    func LastIndexByte(s string, c byte) int
    func LastIndexFunc(s string, f func(rune) bool) int
    func Map(mapping func(rune) rune, s string) string
    func Repeat(s string, count int) string
    func Replace(s, old, new string, n int) string
    func Split(s, sep string) []string
    func SplitAfter(s, sep string) []string
    func SplitAfterN(s, sep string, n int) []string
    func SplitN(s, sep string, n int) []string
    func Title(s string) string
    func ToLower(s string) string
    func ToLowerSpecial(_case unicode.SpecialCase, s string) string
    func ToTitle(s string) string
    func ToTitleSpecial(_case unicode.SpecialCase, s string) string
    func ToUpper(s string) string
    func ToUpperSpecial(_case unicode.SpecialCase, s string) string
    func Trim(s string, cutset string) string
    func TrimFunc(s string, f func(rune) bool) string
    func TrimLeft(s string, cutset string) string
    func TrimLeftFunc(s string, f func(rune) bool) string
    func TrimPrefix(s, prefix string) string
    func TrimRight(s string, cutset string) string
    func TrimRightFunc(s string, f func(rune) bool) string
    func TrimSpace(s string) string
    func TrimSuffix(s, suffix string) string
    type Reader
        func NewReader(s string) *Reader
        func (r *Reader) Len() int
        func (r *Reader) Read(b []byte) (n int, err error)
        func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
        func (r *Reader) ReadByte() (byte, error)
        func (r *Reader) ReadRune() (ch rune, size int, err error)
        func (r *Reader) Reset(s string)
        func (r *Reader) Seek(offset int64, whence int) (int64, error)
        func (r *Reader) Size() int64
        func (r *Reader) UnreadByte() error
        func (r *Reader) UnreadRune() error
        func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
    type Replacer
        func NewReplacer(oldnew ...string) *Replacer
        func (r *Replacer) Replace(s string) string
        func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
