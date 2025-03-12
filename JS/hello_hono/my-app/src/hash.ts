import bcrypt from 'bcrypt';

export function getHash(plainPassword: string): Promise<string> {
  // ソルトラウンド数（計算コスト）
  const saltRounds = '$2b$10$ABCDEFGHIJKLMNOPQRSTUVXYZ12';
  const hashed = bcrypt.hash(plainPassword, saltRounds);
  return hashed;
}
