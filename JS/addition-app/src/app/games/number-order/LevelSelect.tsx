'use client';

import { LevelOption } from '../../../lib/gameUtils';

type Props = {
  onSelectLevel: (level: LevelOption) => void;
};

export default function LevelSelect({ onSelectLevel }: Props) {
  return (
    <div className="flex flex-col items-center justify-center min-h-[80vh] space-y-8 px-4 text-center bg-blue-50">
      <p className="text-2xl font-bold text-gray-800">
        <ruby>レベル<rt>れべる</rt></ruby>を
        <ruby>選<rt>えら</rt></ruby>んでね！
      </p>

      <div className="flex flex-col space-y-6 w-full max-w-xs">
        {(['Lv1', 'Lv2', 'Lv3', 'Lv4'] as LevelOption[]).map((lv) => (
          <button
            key={lv}
            onClick={() => onSelectLevel(lv)}
            className="bg-pink-400 text-white text-xl font-bold px-6 py-6 rounded-full shadow-md hover:bg-pink-500 transition"
          >
            {lv}
          </button>
        ))}
      </div>
    </div>
  );
}
