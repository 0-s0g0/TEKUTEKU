"use client";

import { FormEvent, useState } from "react";
import { useRouter } from "next/navigation";
import styles from "./post.module.css";

export default function PostPage() {
  const router = useRouter();
  const [content, setContent] = useState("");
  const [previewBubble, setPreviewBubble] = useState<string | null>(null);
  const [bubbleAnimating, setBubbleAnimating] = useState(false);
  const [dragOffset, setDragOffset] = useState(0);
  const [pointerStartY, setPointerStartY] = useState<number | null>(null);
  const [hintText, setHintText] = useState("上にスワイプしてバブルを送信");
  const [showCheckmark, setShowCheckmark] = useState(false);
  const [showModal, setShowModal] = useState(false);
  const [school, setSchool] = useState("東京科学大学"); // プルダウンの選択値

  // 送信ボタン押下時、入力内容をバブルプレビューとして表示
  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (content.trim() === "") return;
    setPreviewBubble(content);
    setContent("");
    setDragOffset(0);
    setHintText("上にスワイプしてバブルを送信");
  };

  // pointer（タッチ／マウス）開始時、キャプチャ設定と開始位置（Y座標）の記録
  const handlePointerDown = (e: React.PointerEvent<HTMLDivElement>) => {
    e.currentTarget.setPointerCapture(e.pointerId);
    setPointerStartY(e.clientY);
  };

  // pointer 移動時、Y軸方向のみ移動量を更新（ドラッグ中のみ追従）
  const handlePointerMove = (e: React.PointerEvent<HTMLDivElement>) => {
    if (pointerStartY === null) return;
    const offsetY = e.clientY - pointerStartY;
    setDragOffset(offsetY);
    const safeOffset = offsetY < 0 ? offsetY : 0;
    if (safeOffset < -300) {
      setHintText("指を離して送信");
    } else {
      setHintText("上にスワイプしてバブルを送信");
    }
  };

  // pointer 終了時、キャプチャ解除、上方向に十分移動していれば送信、そうでなければ元の位置に戻す
  const handlePointerUp = (e: React.PointerEvent<HTMLDivElement>) => {
    if (dragOffset < -300) {
      submitBubble();
    } else {
      setDragOffset(0);
      setHintText("上にスワイプしてバブルを送信");
    }
    setPointerStartY(null);
    e.currentTarget.releasePointerCapture(e.pointerId);
  };

  // バブル送信処理：送信時にテキスト（previewBubble）とプルダウンの選択値（school）をPOSTデータに含める
  const submitBubble = async () => {
    try {
      // 実際の送信処理例（API実装に合わせて修正してください）
      // const response = await fetch('/api/submit', {
      //   method: 'POST',
      //   headers: {
      //     'Content-Type': 'application/json',
      //   },
      //   body: JSON.stringify({
      //     text: previewBubble,
      //     schoolId: school,
      //   }),
      // });
      // if (response.ok) {
      //   alert('投稿が成功しました');
      // } else {
      //   alert('投稿に失敗しました');
      // }
    } catch (error) {
      console.error('エラー:', error);
      alert('投稿中にエラーが発生しました');
    }

    setBubbleAnimating(true);
    setHintText("送信済み");
    setTimeout(() => {
      console.log("Bubble submitted:", previewBubble);
      setPreviewBubble(null);
      setBubbleAnimating(false);
      setDragOffset(0);
      setShowCheckmark(true);
      // チェックマークアニメーション終了後にモーダルを表示（1.5秒後）
      setTimeout(() => {
        setShowModal(true);
      }, 1500);
    }, 1000);
  };

  const safeOffset = dragOffset < 0 ? dragOffset : 0;
  const bubbleTransform =
    bubbleAnimating
      ? "translateX(-50%) translateY(-100vh)"
      : pointerStartY !== null
      ? `translateX(-50%) translateY(${safeOffset}px)`
      : "translateX(-50%) translateY(0px)";

  return (
    <div className={styles.container}>
      {/* フロート戻るボタン */}
      <div
        className={styles.backButton}
        onClick={() => router.push("./home")}
      >
        ←
      </div>

      {/* 背景の流れ星エフェクト */}
      <div className={styles.shootingStars}>
        {Array.from({ length: 10 }).map((_, i) => (
          <div key={i} className={styles.shootingStar}></div>
        ))}
      </div>

      {previewBubble && (
        <>
          {/* ヒントテキスト */}
          <div className={styles.hint}>{hintText}</div>
          <div
            className={styles.bubblePreview}
            style={{ transform: bubbleTransform }}
            onPointerDown={handlePointerDown}
            onPointerMove={handlePointerMove}
            onPointerUp={handlePointerUp}
            onPointerCancel={handlePointerUp}
          >
            <div className={styles.bubbleContent}>{previewBubble}</div>
          </div>
        </>
      )}

      <form className={styles.postEditArea} onSubmit={handleSubmit}>
        <select
          className={styles.dropdown}
          value={school}
          onChange={(e) => setSchool(e.target.value)}
        >
          <option value="東京科学大学">東京科学大学</option>
          <option value="九州工業大学">九州工業大学</option>
        </select>
        <textarea
          className={styles.textbox}
          placeholder="バブル投稿を入力してください"
          value={content}
          onChange={(e) => setContent(e.target.value)}
        ></textarea>
        <button type="submit" className={styles.button}>
          送信
        </button>
      </form>

      {/* チェックマークアニメーション */}
      {showCheckmark && (
        <div className={styles.checkmark}>✓</div>
      )}

      {/* モーダル表示 */}
      {showModal && (
        <div className={styles.modal}>
          <div className={styles.modalContent}>
            <p>再び送信しますか？</p>
            <div className={styles.modalButtons}>
              <button
                className={`${styles.modalButton} ${styles.yes}`}
                onClick={() => {
                  setShowModal(false);
                  setShowCheckmark(false);
                }}
              >
                はい
              </button>
              <button
                className={`${styles.modalButton} ${styles.no}`}
                onClick={() => {
                  router.push("pages/home");
                }}
              >
                いいえ
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}