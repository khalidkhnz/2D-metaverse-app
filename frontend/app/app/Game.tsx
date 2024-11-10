"use client";

import { useEffect, useRef } from "react";
import * as Phaser from "phaser";
import useWindowDimensions from "@/hooks/useWindowDimensions";

const Game: React.FC = () => {
  const gameRef = useRef<Phaser.Game | null>(null);
  const { height: windowHeight, width: windowWidth } = useWindowDimensions();

  useEffect(() => {
    if (gameRef.current) return;

    // Main game scene with type definitions
    class MainScene extends Phaser.Scene {
      ball!: Phaser.GameObjects.Arc;
      cursors!: Phaser.Types.Input.Keyboard.CursorKeys;

      constructor() {
        super({ key: "MainScene" });
      }

      preload() {
        // Preload assets if needed
      }

      create() {
        // Create a circle ball at the center of the screen
        this.ball = this.add.circle(400, 300, 20, 0x00ff00); // Green circle

        // Enable physics for the ball
        this.physics.add.existing(this.ball);
        (this.ball.body as Phaser.Physics.Arcade.Body).setCollideWorldBounds(
          true,
        );

        // Set up arrow keys for movement
        this.cursors =
          this.input.keyboard?.createCursorKeys() as Phaser.Types.Input.Keyboard.CursorKeys;
      }

      update() {
        const speed = 200;

        // Reset velocity
        const ballBody = this.ball.body as Phaser.Physics.Arcade.Body;
        ballBody.setVelocity(0);

        // Horizontal movement
        if (this.cursors.left?.isDown) {
          ballBody.setVelocityX(-speed);
        } else if (this.cursors.right?.isDown) {
          ballBody.setVelocityX(speed);
        }

        // Vertical movement
        if (this.cursors.up?.isDown) {
          ballBody.setVelocityY(-speed);
        } else if (this.cursors.down?.isDown) {
          ballBody.setVelocityY(speed);
        }
      }
    }

    const config: Phaser.Types.Core.GameConfig = {
      type: Phaser.AUTO,
      width: windowWidth,
      height: windowHeight,
      mode: Phaser.Scale.FIT,
      autoCenter: Phaser.Scale.CENTER_BOTH,
      physics: {
        default: "arcade",
        arcade: {
          gravity: { x: 0, y: 0 },
          debug: false,
        },
      },
      scene: MainScene,
    };

    // Initialize the Phaser game
    gameRef.current = new Phaser.Game(config);

    // Cleanup Phaser instance on component unmount
    return () => {
      if (gameRef.current) {
        gameRef.current.destroy(true);
        gameRef.current = null;
      }
    };
  }, [windowHeight, windowWidth]);

  return <div id="phaser-game" />;
};

export default Game;
