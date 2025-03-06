"use client";

import Link from "next/link";
import styles from "./home.module.css";

export default function HomePage() {
  return (
    <div className={styles.container}>
      <h1 className={styles.title}>TekuTeku - HOME</h1>
      <div className={styles.buttonGroup}>
        <Link href="./post">
          <button className={styles.button}>投稿する</button>
        </Link>
        <Link href="./index">
          <button className={styles.button}>返信する</button>
        </Link>
        <Link href="./index">
          <button className={styles.button}>大画面表示</button>
        </Link>
      </div>
    </div>
  );
}