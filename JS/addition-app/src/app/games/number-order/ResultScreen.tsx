'use client';

type Props = {
  startTime: number;
  endTime: number;
  onRetry: () => void;
  onBack: () => void;
};

export default function ResultScreen({ startTime, endTime, onRetry, onBack }: Props) {
  const timeTaken = ((endTime - startTime) / 1000).toFixed(2);

  return (
    <div className="flex flex-col items-center justify-center h-[80vh] space-y-6">
      <h2 className="text-2xl font-bold text-green-600">🎉 クリア！</h2>
      <p className="text-lg">かかった時間：{timeTaken} 秒</p>
      <div className="flex space-x-4">
        <button
          onClick={onRetry}
          className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
        >
          もう一度
        </button>
        <button
          onClick={onBack}
          className="bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600"
        >
          レベル選択に戻る
        </button>
      </div>
    </div>
  );
}