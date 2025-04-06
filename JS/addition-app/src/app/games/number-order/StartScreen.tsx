// app/game/StartScreen.tsx
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
    <div className="flex flex-col items-center mt-10 space-y-4">
      <p className="text-lg font-semibold">レベル：{level}</p>
      <button
        onClick={handleStart}
        className="bg-green-500 text-white px-6 py-3 rounded-lg text-lg hover:bg-green-600"
      >
        スタート！
      </button>
    </div>
  );
}