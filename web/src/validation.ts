declare interface Validation {}

export interface IPasswordValidation extends Validation {
  hasMinLength: boolean;
  hasLowerCase: boolean;
  hasUpperCase: boolean;
  hasNumber: boolean;
  hasSpecialCharacter: boolean;
}

export interface IEmailValidation extends Validation {
  isNotEmpty: boolean;
  isValidEmail: boolean;
}

export interface IConfirmPasswordValidation extends Validation {
  isMatch: boolean;
}

export const isValid = (v: Validation) => {
  const arr = Object.values(v) as boolean[];
  return !arr.includes(false);
};

export const validatePassword = (password: string) => {
  // Assume that user passwords will be sent to a server to be hashed.
  // All validation measures exist to increase entropy, ie prevent
  // some bozo from guessing the password
  // Special characters should not be limited for the sake of sanitization
  // Don't set a max length, it may convey a security risk to the end user
  // Sanitzation and handling
  const MINIMUM_PASSWORD_LENGTH = 8;

  const validation: IPasswordValidation = {
    hasMinLength: false,
    hasLowerCase: false,
    hasUpperCase: false,
    hasNumber: false,
    hasSpecialCharacter: false,
  };

  if (password.length >= MINIMUM_PASSWORD_LENGTH) {
    validation.hasMinLength = true;
  }

  for (let i = 0; i < password.length; i++) {
    if (!validation.hasLowerCase && /[a-z]/.test(password[i])) {
      validation.hasLowerCase = true;
    }
    if (!validation.hasUpperCase && /[A-Z]/.test(password[i])) {
      validation.hasUpperCase = true;
    }
    if (!validation.hasNumber && /[\d]/.test(password[i])) {
      validation.hasNumber = true;
    }
    if (
      !validation.hasSpecialCharacter &&
      !/[a-z]/.test(password[i]) &&
      !/[A-Z]/.test(password[i]) &&
      !/[\d]/.test(password[i])
    ) {
      validation.hasSpecialCharacter = true;
    }
  }

  return validation;
};

export const validateEmail = (email: string) => {
  const validation: IEmailValidation = {
    isNotEmpty: false,
    isValidEmail: false,
  };
  if (email !== "") {
    validation.isNotEmpty = true;
  }
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  const consecutiveDotRegex = /\.{2,}/;
  if (emailRegex.test(email) && !consecutiveDotRegex.test(email)) {
    validation.isValidEmail = true;
  }
  return validation;
};

export const validateConfirmPassword = (
  password: string,
  confirmPassword: string
) => {
  const validation: IConfirmPasswordValidation = {
    isMatch: false,
  };

  if (password === confirmPassword) {
    validation.isMatch = true;
  }
  return validation;
};
