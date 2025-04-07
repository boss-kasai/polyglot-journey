// app/game/ResultScreen.tsx
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
    <div className="flex flex-col items-center justify-center h-[80vh] space-y-8 px-4 text-center bg-green-50">
      <h2 className="text-3xl font-bold text-green-600">🎉 クリア！</h2>
      <p className="text-xl text-gray-700 font-semibold">
        かかった<ruby>時間<rt>じかん</rt></ruby>：<span className="text-blue-600">{timeTaken} 秒</span>
      </p>

      <div className="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-6 w-full max-w-xs">
        <button
          onClick={onRetry}
          className="bg-blue-400 text-white text-xl font-bold px-6 py-4 rounded-full shadow-md hover:bg-blue-500 transition whitespace-nowrap"
        >
          🔁 <ruby>もう一度<rt>いちど</rt></ruby>
        </button>

        <button
          onClick={onBack}
          className="bg-gray-400 text-white text-xl font-bold px-6 py-4 rounded-full shadow-md hover:bg-gray-500 transition whitespace-nowrap"
        >
          🔙 <ruby>レベル<rt>れべる</rt></ruby>
          <ruby>選択<rt>せんたく</rt></ruby>に
          <ruby>戻<rt>もど</rt></ruby>る
        </button>
      </div>
    </div>
  );
}
