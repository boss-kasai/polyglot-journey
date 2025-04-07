'use client';

import { LevelOption, generateNumbers } from '@/lib/gameUtils';

type Props = {
  level: LevelOption;
  onStart: (numbers: { value: number; x: number; y: number }[]) => void;
};

export default function StartScreen({ level, onStart }: Props) {
  const handleStart = () => {
    const generated = generateNumbers(level);
    onStart(generated);
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-[80vh] space-y-8 px-4 text-center bg-green-50">
      <p className="text-2xl font-bold text-green-700">
        <ruby>レベル<rt>れべる</rt></ruby>：{level}
      </p>

      <button
        onClick={handleStart}
        className="bg-green-400 text-white text-xl font-bold px-8 py-6 rounded-full shadow-md hover:bg-green-500 transition"
      >
        <ruby>スタート<rt>すたーと</rt></ruby>！
      </button>
    </div>
  );
}
