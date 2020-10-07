package main

import (
	"fmt"
	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
)

func Stop(gopigo3 *g.Driver) {
	err := gopigo3.SetMotorDps(g.MOTOR_LEFT+g.MOTOR_RIGHT, 0)
	if err != nil {
		fmt.Errorf("Error stopping the robot: %+v", err)
	}
}

func Forward(gopigo3 *g.Driver, speed int) {
	err := gopigo3.SetMotorDps(g.MOTOR_LEFT+g.MOTOR_RIGHT, speed)
	if err != nil {
		fmt.Errorf("Error moving the robot forward: %+v", err)
	}
}

func Backward(gopigo3 *g.Driver, speed int) {
	err := gopigo3.SetMotorDps(g.MOTOR_LEFT+g.MOTOR_RIGHT, speed*-1)
	if err != nil {
		fmt.Errorf("Error moving the robot backward: %+v", err)
	}
}

func Right(gopigo3 *g.Driver, speed int) {
	err := gopigo3.SetMotorDps(g.MOTOR_LEFT, speed)
	if err != nil {
		fmt.Errorf("Error moving the robot backward: %+v", err)
	}
	err = gopigo3.SetMotorDps(g.MOTOR_RIGHT, 0)
	if err != nil {
		fmt.Errorf("Error moving the robot backward: %+v", err)
	}
}

func SpinRight(gopigo3 *g.Driver, speed int) {
	err := gopigo3.SetMotorDps(g.MOTOR_LEFT, speed)
	if err != nil {
		fmt.Errorf("Error moving the robot backward: %+v", err)
	}
	err = gopigo3.SetMotorDps(g.MOTOR_RIGHT, speed*-1)
	if err != nil {
		fmt.Errorf("Error moving the robot backward: %+v", err)
	}
}

func Left(gopigo3 *g.Driver, speed int) {
	err := gopigo3.SetMotorDps(g.MOTOR_LEFT, 0)
	if err != nil {
		fmt.Errorf("Error moving the robot backward: %+v", err)
	}
	err = gopigo3.SetMotorDps(g.MOTOR_RIGHT, speed)
	if err != nil {
		fmt.Errorf("Error moving the robot backward: %+v", err)
	}
}

func SpinLeft(gopigo3 *g.Driver, speed int) {
	err := gopigo3.SetMotorDps(g.MOTOR_LEFT, speed*-1)
	if err != nil {
		fmt.Errorf("Error moving the robot backward: %+v", err)
	}
	err = gopigo3.SetMotorDps(g.MOTOR_RIGHT, speed)
	if err != nil {
		fmt.Errorf("Error moving the robot backward: %+v", err)
	}
}
