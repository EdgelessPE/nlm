import { ZSTDDecoder } from "zstddec";

export function formatFileSize(size: number) {
  if (size < 1024) {
    return `${size} B`;
  }
  if (size < 1024 * 1024) {
    return `${(size / 1024).toFixed(2)} KB`;
  }
  if (size < 1024 * 1024 * 1024) {
    return `${(size / 1024 / 1024).toFixed(2)} MB`;
  }
  return `${(size / 1024 / 1024 / 1024).toFixed(2)} GB`;
}

export function base64Decode(str: string) {
  const binaryString = atob(str);

  const byteArray = new Uint8Array(binaryString.length);
  for (let i = 0; i < binaryString.length; i++) {
    byteArray[i] = binaryString.charCodeAt(i);
  }

  const decoder = new TextDecoder("utf-8");
  return decoder.decode(byteArray);
}

export async function decodeZstd(raw: Uint8Array, size: number) {
  const decoder = new ZSTDDecoder();
  await decoder.init();
  return decoder.decode(raw, size);
}
