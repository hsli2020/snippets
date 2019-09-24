package main // wc-naive

import (
	"bufio"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("no file path specified")
	}
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	const bufferSize = 16 * 1024
	reader := bufio.NewReaderSize(file, bufferSize)

	lineCount := 0
	wordCount := 0
	byteCount := 0

	prevByteIsSpace := true
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		byteCount++

		switch b {
		case '\n':
			lineCount++
			prevByteIsSpace = true
		case ' ', '\t', '\r', '\v', '\f':
			prevByteIsSpace = true
		default:
			if prevByteIsSpace {
				wordCount++
				prevByteIsSpace = false
			}
		}
	}

	println(lineCount, wordCount, byteCount, file.Name())
}

package main // wc-chunks

import (
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("no file path specified")
	}
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalCount := Count{}
	lastCharIsSpace := true

	const bufferSize = 16 * 1024
	buffer := make([]byte, bufferSize)

	for {
		bytes, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		count := GetCount(Chunk{lastCharIsSpace, buffer[:bytes]})
		lastCharIsSpace = IsSpace(buffer[bytes-1])

		totalCount.LineCount += count.LineCount
		totalCount.WordCount += count.WordCount
	}

	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	byteCount := fileStat.Size()

	println(totalCount.LineCount, totalCount.WordCount, byteCount, file.Name())
}

type Chunk struct {
	PrevCharIsSpace bool
	Buffer          []byte
}

type Count struct {
	LineCount int
	WordCount int
}

func GetCount(chunk Chunk) Count {
	count := Count{}

	prevCharIsSpace := chunk.PrevCharIsSpace
	for _, b := range chunk.Buffer {
		switch b {
		case '\n':
			count.LineCount++
			prevCharIsSpace = true
		case ' ', '\t', '\r', '\v', '\f':
			prevCharIsSpace = true
		default:
			if prevCharIsSpace {
				prevCharIsSpace = false
				count.WordCount++
			}
		}
	}

	return count
}

func IsSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r' || b == '\v' || b == '\f'
}


package main // wc-channel

import (
	"io"
	"os"
	"runtime"
)

func main() {
	if len(os.Args) < 2 {
		panic("no file path specified")
	}
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	chunks := make(chan Chunk)
	counts := make(chan Count)

	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		go ChunkCounter(chunks, counts)
	}

	const bufferSize = 16 * 1024
	lastCharIsSpace := true

	for {
		buffer := make([]byte, bufferSize)
		bytes, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		chunks <- Chunk{lastCharIsSpace, buffer[:bytes]}
		lastCharIsSpace = IsSpace(buffer[bytes-1])
	}
	close(chunks)

	totalCount := Count{}
	for i := 0; i < numWorkers; i++ {
		count := <-counts
		totalCount.LineCount += count.LineCount
		totalCount.WordCount += count.WordCount
	}
	close(counts)

	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	byteCount := fileStat.Size()

	println("%d %d %d %s\n", totalCount.LineCount, totalCount.WordCount, byteCount, file.Name())
}

func ChunkCounter(chunks <-chan Chunk, counts chan<- Count) {
	totalCount := Count{}
	for {
		chunk, ok := <-chunks
		if !ok {
			break
		}
		count := GetCount(chunk)
		totalCount.LineCount += count.LineCount
		totalCount.WordCount += count.WordCount
	}
	counts <- totalCount
}

type Chunk struct {
	PrevCharIsSpace bool
	Buffer          []byte
}

type Count struct {
	LineCount int
	WordCount int
}

func GetCount(chunk Chunk) Count {
	count := Count{}

	prevCharIsSpace := chunk.PrevCharIsSpace
	for _, b := range chunk.Buffer {
		switch b {
		case '\n':
			count.LineCount++
			prevCharIsSpace = true
		case ' ', '\t', '\r', '\v', '\f':
			prevCharIsSpace = true
		default:
			if prevCharIsSpace {
				prevCharIsSpace = false
				count.WordCount++
			}
		}
	}

	return count
}

func IsSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r' || b == '\v' || b == '\f'
}

package main // wc-mutex

import (
	"io"
	"os"
	"runtime"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		panic("no file path specified")
	}
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileReader := &FileReader{
		File:            file,
		LastCharIsSpace: true,
	}
	counts := make(chan Count)

	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		go FileReaderCounter(fileReader, counts)
	}

	totalCount := Count{}
	for i := 0; i < numWorkers; i++ {
		count := <-counts
		totalCount.LineCount += count.LineCount
		totalCount.WordCount += count.WordCount
	}
	close(counts)

	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	byteCount := fileStat.Size()

	println(totalCount.LineCount, totalCount.WordCount, byteCount, file.Name())
}

type FileReader struct {
	File            *os.File
	LastCharIsSpace bool
	mutex           sync.Mutex
}

func (fileReader *FileReader) ReadChunk(buffer []byte) (Chunk, error) {
	fileReader.mutex.Lock()
	defer fileReader.mutex.Unlock()

	bytes, err := fileReader.File.Read(buffer)
	if err != nil {
		return Chunk{}, err
	}

	chunk := Chunk{fileReader.LastCharIsSpace, buffer[:bytes]}
	fileReader.LastCharIsSpace = IsSpace(buffer[bytes-1])

	return chunk, nil
}

func FileReaderCounter(fileReader *FileReader, counts chan Count) {
	const bufferSize = 16 * 1024
	buffer := make([]byte, bufferSize)

	totalCount := Count{}

	for {
		chunk, err := fileReader.ReadChunk(buffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		count := GetCount(chunk)
		totalCount.LineCount += count.LineCount
		totalCount.WordCount += count.WordCount
	}

	counts <- totalCount
}

type Chunk struct {
	PrevCharIsSpace bool
	Buffer          []byte
}

type Count struct {
	LineCount int
	WordCount int
}

func GetCount(chunk Chunk) Count {
	count := Count{}

	prevCharIsSpace := chunk.PrevCharIsSpace
	for _, b := range chunk.Buffer {
		switch b {
		case '\n':
			count.LineCount++
			prevCharIsSpace = true
		case ' ', '\t', '\r', '\v', '\f':
			prevCharIsSpace = true
		default:
			if prevCharIsSpace {
				prevCharIsSpace = false
				count.WordCount++
			}
		}
	}

	return count
}

func IsSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r' || b == '\v' || b == '\f'
}
