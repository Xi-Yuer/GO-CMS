import SparkMD5 from 'spark-md5';
import { mergeFileChunk, uploadFileChunk } from '@/service/api/file';

interface Chunk {
  file: Blob;
  index: number;
}

/**
 * 大文件上传
 * @param file 文件对象
 * @param fileHash
 * @param fallBack
 */
export const BigFileUpload = async (file: File, fileHash: string, fallBack?: (size: number) => void) => {
  // 2. 将文件分片
  const chunkSize = 1024 * 1024; // 每个分片大小为1MB
  const chunks = createFileChunks(file, chunkSize);

  // 3. 上传分片
  for (let i = 0; i < chunks.length; i++) {
    // 上传分片
    await uploadFileChunk({
      file: chunks[i].file,
      identifier: fileHash,
    });
    fallBack && fallBack(Math.ceil(((i + 1) * 100) / chunks.length));
  }

  // 4. 通知服务器合并分片
  await mergeFileChunk({ identifier: fileHash, fileName: file.name, fileSize: file.size });
};

/**
 * 计算文件哈希
 * @param file 文件对象
 * @returns 文件的哈希值
 */
export const calculateFileHash = (file: File): Promise<string> => {
  return new Promise((resolve, reject) => {
    const chunkSize = 2 * 1024 * 1024; // 每个分片大小为2MB
    const chunks = Math.ceil(file.size / chunkSize);
    const spark = new SparkMD5.ArrayBuffer();
    const fileReader = new FileReader();

    let currentChunk = 0;

    fileReader.onload = (e) => {
      if (e.target?.result) {
        spark.append(e.target.result as ArrayBuffer);
        currentChunk++;

        if (currentChunk < chunks) {
          loadNext();
        } else {
          const hash = spark.end();
          resolve(hash);
        }
      } else {
        reject('File read error');
      }
    };

    fileReader.onerror = () => {
      reject('Error reading file');
    };

    const loadNext = () => {
      const start = currentChunk * chunkSize;
      const end = start + chunkSize >= file.size ? file.size : start + chunkSize;
      fileReader.readAsArrayBuffer(file.slice(start, end));
    };

    loadNext();
  });
};

/**
 * 创建文件分片
 * @param file 文件对象
 * @param chunkSize 分片大小
 * @returns 分片数组
 */
export const createFileChunks = (file: File, chunkSize: number): Chunk[] => {
  const chunks: Chunk[] = [];
  let currentChunk = 0;
  while (currentChunk < file.size) {
    chunks.push({
      file: file.slice(currentChunk, currentChunk + chunkSize),
      index: chunks.length,
    });
    currentChunk += chunkSize;
  }
  return chunks;
};
